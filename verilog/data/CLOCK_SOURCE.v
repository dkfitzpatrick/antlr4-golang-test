//  ----------------------------------------------------------------------------
//                    Copyright Message
//  ----------------------------------------------------------------------------
//
//  CONFIDENTIAL and PROPRIETARY
//  COPYRIGHT (c) NXP B.V. 2015-2016
//
//  All rights are reserved. Reproduction in whole or in part is
//  prohibited without the written consent of the copyright owner.
//
//  CoReUse 4.5 information header.
//
//  ----------------------------------------------------------------------------
//                    Design Information
//  ----------------------------------------------------------------------------
//  File            : sync_support.v
//  Organisation    : MCO
//  Tag             : 1.1.9_a25
//  Date            : $Date: Wed Jul 31 17:01:43 2019 $
//  Revision        : $Revision: 1.60.1.1 $
//
//  IP Name         : SYNC_ modules for general use
//  Description     : Clock crossing support for Slave
//    This contains support for cross-clock-domain (CDC) synchronization
//    of type 2-phase and other, covering pulses and levels. Used for 
//    SCL to (system) CLK and CLK to SCL. Used for single nets, handshakes,
//    FIFOs, etc. 
//    Note use of 2-phase handshakes due to SCL stopping suddenly.
//
//  ----------------------------------------------------------------------------
//                    Revision History
//  ----------------------------------------------------------------------------
//
//
//  ----------------------------------------------------------------------------
//                    Implementation details
//  ----------------------------------------------------------------------------
//  See the micro-arch spec and MIPI I3C spec
//
//    This suppports the Clock Domain Crossing using named blocks
//    so Spyglass (or equiv) can be told that the sync is here.
//    For most modern process, the single flop model is fine, although
//    the flop type may need to be changed by constraint for older
//    process (e.g. a sync flop, which has a "gravity" to settle
//    faster). 
//    1 flop is fine in normal cases because the Q out metastable
//    noise will be short enough in time for the paths used at the 
//    lower speeds. That is, there is ~40ns minimum to both settle
//    and arrive settled at the other end. 
//  ----------------------------------------------------------------------------

// Note naming: SYNC_= synchronizing, 2PH_=2-phase, S2P_=SCL to CLK, STATE=state with clr in
//              LVL_=level vs. pulse input, LVLH_=level high
// Other naming: ASet=async set, local clear, AClr=local set, async clear
//               Seq2=2 flop sequencer to ensure we get 1 pulse in local domain
module SYNC_2PH_S2C_STATE( 
  input             rst_n,
  input             scl,
  input             clk,
  input             trig_scl,   // trigger input from SCL domain
  output            out_clk,    // output state in CLK domain
  input             clear_clk   // input clear from CLK domain
  );

  reg               scl_hshake;
  reg               clk_state;
  reg               clk_ack;

  // NOTE: this can be a bit confusing. The scl_  to clk_
  //       is 2-phase. The scl_hshake changes on the trigger.
  //       The clk_state changes on hshake^ack==1 for the 
  //       purpose of the use in the CLK domain. The clk_ack
  //       only changes when the clk_state has, so that we
  //       do not miss it and yet we do not expose metastable
  //       logic (hshake^ack) outside.
  //       We use 2-phase because SCL is not free running and
  //       may stop. So, CLK does all of the work.
  //       Note that state based ones are only cleared by
  //       explit request; so more changes may happen before
  //       cleared and that is OK.

  always @ (posedge scl or negedge rst_n)
    if (!rst_n)
      scl_hshake <= 1'b0;
    else if (trig_scl)                  // trigger pulse (1 clock in scl)
      scl_hshake <= ~scl_hshake;        // flips state

  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_state  <= 1'b0;
    else if (scl_hshake ^ clk_ack)      // CDC
      clk_state  <= 1'b1;               // sets on diff
    else if (clear_clk)
      clk_state  <= 1'b0;               // clears on explicit clear
  assign out_clk = clk_state;

  // ACK only chnages after state does and if still diff
  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_ack    <= 1'b0;
    else if (clk_state & (scl_hshake^clk_ack)) // change on state - CDC
      clk_ack    <= ~clk_ack; 

endmodule 

module SYNC_2PH_LVL_S2C_STATE(
  input             rst_n,
  input             scl,
  input             clk,
  input             trig_scl,   // trigger level from SCL domain
  output            out_clk,    // output state in CLK domain
  input             clear_clk   // input clear from CLK domain
  );

  reg               clk_state;
  reg               clk_ack;

  // see SYNC_2PH_S2C_STATE for details. Only difference is
  // level held does not need scl_hshake since level (and 
  // if level too short for clk, then not registered). If
  // we want to make sure seen, then need independent bit
  // FUTURE: look at that for cases of stopped CLK

  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_state   <= 1'b0;
    else if (trig_scl ^ clk_ack) // level change - CDC
      clk_state   <= 1'b1; 
    else if (clear_clk)
      clk_state   <= 1'b0; 
  assign out_clk = clk_state;

  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_ack     <= 1'b0;
    else if (clk_state &
            (trig_scl ^ clk_ack)) // still diff - CDC
      clk_ack     <= ~clk_ack; 

endmodule

module SYNC_2PH_LVLH_S2C_STATE(
  input             rst_n,
  input             scl,
  input             clk,
  input             trig_scl,   // trigger level from SCL domain
  output            out_clk,    // output state in CLK domain
  input             clear_clk   // input clear from CLK domain
  );

  reg               clk_state;
  reg         [1:0] clk_ack;

  // see SYNC_2PH_S2C_STATE for details. 
  // Level High is a bit tricky. Clk_state only
  // goes high if level becomes High (edge if you
  // will). So, we need to remember what state it
  // was in so we only catch the edge. But, because
  // this across the clock domain, we cannot just
  // compare with the SCL level, since we may see 
  // it sooner or later and so miss the edge.
  // So we use an extra flop
  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_ack     <= 2'b00;
    else if (clk_ack[0] ^ trig_scl)
      clk_ack[0]  <= ~clk_ack[0];       // CDC
    else if (^clk_ack)
      clk_ack[1]  <= ~clk_ack[1];       // 1 cycle behind

  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_state   <= 1'b0;
    else if (^clk_ack & ~clk_ack[1])
      clk_state   <= 1'b1; 
    else if (clear_clk)
      clk_state   <= 1'b0; 
  assign out_clk = clk_state;

endmodule

// Next is set async and clear local
module SYNC_Pulse_S2C(
  input SCL,
  input CLK,
  input RSTn,
  input local_set,
  output o_pulse);

  // 4 phase handshake, but pulse out
  reg  svalue, cvalue, cpulse;
  always @ (posedge SCL or negedge RSTn)
    if (!RSTn)
      svalue <= 1'b0;
    else if (local_set)
      svalue <= 1'b1;
    else if (cvalue) // CDC
      svalue <= 1'b0;

  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      cvalue <= 1'b0;
    else if (svalue) // CDC
      cvalue <= 1'b1;
    else 
      cvalue <= 1'b0;
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      cpulse <= 1'b0;
    else if (cvalue)
      cpulse <= 1'b1;
    else
      cpulse <= 1'b0;
  assign o_pulse = cvalue & ~cpulse;

endmodule


// Next is set async and clear local, with sequence for 1 pulse out
module SYNC_ASet_Seq2(
  input CLK,
  input RSTn,
  input async_set,
  input local_clear, 
  output o_pulse);

  reg  value, seq;
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      value <= 1'b0;
    else if (async_set) // CDC
      value <= 1'b1;
    else if (local_clear) 
      value <= 1'b0;
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      seq <= 1'b0;
    else if (seq ^ value)
      seq <= ~seq;
  // below pulses for 1 clock high, else is low
  assign o_pulse = value & ~seq;

endmodule

// Next is set local and clear async,  with sequence for 1 pulse out
module SYNC_AClr_Seq2(
  input CLK,
  input RSTn,
  input local_set,
  input async_clear, 
  output o_pulse);

  reg  value, seq;
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      value <= 1'b0;
    else if (local_set) 
      value <= 1'b1;
    else if (async_clear) // CDC
      value <= 1'b0;
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      seq <= 1'b0;
    else if (seq ^ value)
      seq <= ~seq;
  // below pulses for 1 clock high, else is low
  assign o_pulse = value & ~seq;

endmodule

// next is straight sync from one domain to the other
// for local use - 1 flop
module SYNC_S2C #(parameter WIDTH=1) ( 
  input             rst_n,
  input             clk,
  input  [WIDTH-1:0]scl_data,
  output [WIDTH-1:0]out_clk     // output copy in CLK domain
  );

  reg  [WIDTH-1:0]  clk_copy;

  assign out_clk = clk_copy;

  // note: could use clk_copy^scl_data as test to allow ICG
  always @ (posedge clk or negedge rst_n)
    if (!rst_n)
      clk_copy <= {WIDTH{1'b0}};
    else 
      clk_copy <= scl_data;
endmodule

// next is straight sync from one domain to the other
// for local use - 1 flop
module SYNC_C2S #(parameter WIDTH=1) ( 
  input             rst_n,
  input             scl,        // could be SCL_n as well
  input  [WIDTH-1:0]clk_data,
  output [WIDTH-1:0]out_scl     // output copy in CLK domain
  );

  reg   [WIDTH-1:0] scl_copy;

  assign out_scl = scl_copy;

  // note: could use clk_copy^scl_data as test to allow ICG
  always @ (posedge scl or negedge rst_n)
    if (!rst_n)
      scl_copy <= {WIDTH{1'b0}};
    else 
      scl_copy <= clk_data;
endmodule

// Next is set local and clear async from CLK to SCL
module SYNC_AClr_C2S(
  input CLK, // system domain
  input RSTn,
  input local_set,
  input async_clear,
  output o_value);

  reg  value;                   // CLK domain
  always @ (posedge CLK or negedge RSTn)
    if (!RSTn)
      value <= 1'b0;
    else if (local_set) 
      value <= 1'b1;
    else if (async_clear) // CDC
      value <= 1'b0;
  assign o_value = value;

endmodule

//  ----------------------------------------------------------------------------
//                    Copyright Message
//  ----------------------------------------------------------------------------
//
//  CONFIDENTIAL and PROPRIETARY
//  COPYRIGHT (c) NXP B.V. 2015-2020
//
//  All rights are reserved. Reproduction in whole or in prt is
//  prohibited without the written consent of the copyright owner.
//
//  CoReUse 4.5 information header.
//
//  ----------------------------------------------------------------------------
//                    Design Information
//  ----------------------------------------------------------------------------
//  File            : i3c_sdr_slave_engine.v
//  Organisation    : MCO
//  Tag             : 1.1.9_a25
//  Date            : $Date: Tue Nov 10 09:02:53 2020 $
//  Revision        : $Revision: 1.60.1.2.1.12 $
//
//  IP Name         : i3c_sdr_slave_engine 
//  Description     : MIPI I3C and I2C Slave processing main component
//    This is the SDR engine which also handles legacy i2c as well. This 
//    processes START and STOP, the headers, matches various addresses, 
//    serializes and deserializes the data bytes, performas parity checking
//    where needed, and can be held in suppressed state (for HDR and so on).
//    See also: exit pattern detector, ENTDAA/SETDASA, and CCC handling.
//    Note that HDR-DDR uses parts of this for its handling.
//
//  ----------------------------------------------------------------------------
//                    Revision History
//  ----------------------------------------------------------------------------
//
//
//  ----------------------------------------------------------------------------
//                    Implementation details
//  ----------------------------------------------------------------------------
//  See the micro-arch spec and MIPI I3C spec
//
//  -- Clocks and Resets --
//    There are 3 separate clock domains fed in from above, although all related. 
//    We have SCL, SCL_n (inverted SCL), and SDA. When in scan test, they are
//    all the same (so we use all posedge flops).
//    The inversion and muxing is handled above this layer to be sure it is
//    glitchfree and clean; it usually means use of special cells or macros,
//    which also allows naming the clock sources and clock relationships for
//    STA and timing closure.
//    The Clocks are:
//    1. clk_SCL_n for falling edge SCL, which is the most widely used.
//       -- We launch changes in SDA on this edge
//       -- We detect special cases like repeated START and START
//       -- State machine advances on this edge
//       -- 'Use' of registered SDA is on this edge
//    2. clk_SCL for rising edge SCL, which is used in fewer places
//       -- We sample and register SDA from Master on this edge. 
//       -- We tri-state for read T bit on this edge, but that is combinatorial,
//          so not really in SCL domain (using SCL level instead).
//    3. clk_SDA for rising edge SDA, which detects STOP (when SCL=1). 
//       -- This is one flop. It is set when SCL is High. This is how
//           we know a START is a real START and not a repeated START
//       -- This is reset by normal reset and also real_START_d (registered START). 
//    There are 2 resets, a hold, and a clock gating control. These are:
//    1. Global reset which is fed from system and is normally POR.
//    2. Local reset, which is only used related to START/STOP.
//    3. Hold, which is used when in HDR.
//    4. Clock gating, which is used to save power.
//  -- Data buffering for From-Bus --
//    The block owns 16b of flops for buffering. In SDR, this is used as 
//    a ping pong buffer model of 2x8b. The outer layer manages the 
//    buffer use. 
//    The fb_data_use input controls which data buffer to use. If neither
//    is selected, that will cause an overrun if needed.
//    The block signals with a 1 SCL_n pulse when the current buffer is done.
//    The outer layer uses SCL (half edge) to sample/change its nets and
//    so, we avoid timing races. This presumes that ~38ns is more than enough
//    for these narrow cones and routes (between engine and upper layer).
//  -- Data buffering for To-Bus --
//    The block gets data from the upper layer to return on Read using a ping
//    pong buffering model with a 2 bit valid flag indicating which to use. 
//    If none is available at the RnW point, the address is NACKed and the
//    tb_urun_ack output is pulsed.
//    Once in Reasd, if none is available by the 9th bit, it terminates
//    the Read and pulses the tb_urun_nack output.
//    If there is data, it is processed and the tb_data_ack output is 
//    pulsed when done. 
//  -- Status and interrupts --
//    The block pulses various states and interrupts (events) as they happen,
//    such as start, address matched, event sent, STOP (which is held while
//    true).
//    It also tracks critical states/modes related to CCCs. This is done
//    whether the block is processing CCCs or the SW is. The low level
//    engine has to know the rules for broadcast/direct and when that 
//    state ends, DAA mode, and HDR mode. 
//    DAA mode is supported in service to the upper layer handling as a
//    separate detail.
//  -- IBI/MR/HJ --
//    The upper layer signals events, with priority implicitly as HJ (if
//    Hot Join used), IBI (if interrupt used), and MR (if Master request
//    or Peer-to-Peer request).
//    Once pending, the block will wait for a START (not repeated START)
//    and try to win on arbitration. It will report if it won and if ACKed
//    (vs. NACKed) by Master, leaving it to Upper layer to handle, except
//    for IBI with byte following, which block will emit if used.
//    If the Bus is Available (Bus Free condition for a certain period),
//    then the upper layer can request a forced START. This is handled
//    combinatorially but only when STOP_r is true and SCL=1.
//  ----------------------------------------------------------------------------

//  ----------------------------------------------------------------------------
//                    Copyright Message
//  ----------------------------------------------------------------------------
//
//  CONFIDENTIAL and PROPRIETARY
//  COPYRIGHT (c) NXP B.V. 2015-2020
//
//  All rights are reserved. Reproduction in whole or in part is
//  prohibited without the written consent of the copyright owner.
//
//  CoReUse 4.5 information header.
//
//  ----------------------------------------------------------------------------
//                    Design Information
//  ----------------------------------------------------------------------------
//  File            : i3c_params.v
//  Organisation    : MCO
//  Tag             : 1.1.9_a25
//  Date            : $Date: Tue May  5 21:20:48 2020 $
//  Revision        : $Revision: 1.60.1.1.1.3 $
//
//  IP Name         : i3c_params 
//  Description     : MIPI I3C related parameters/constants
//    This contains the defined params as named constants. Note that this
//    is not related to the parametrization of the block per se, but does
//    have the constants used by parameters including the bit breakouts.
//
//  ----------------------------------------------------------------------------
//                    Revision History
//  ----------------------------------------------------------------------------
//
//
//  ----------------------------------------------------------------------------
//                    Implementation details
//  ----------------------------------------------------------------------------
//  See the micro-arch spec and MIPI I3C spec
//
//    Note the use of `define vs. localparam
//  ----------------------------------------------------------------------------

  //
  // Instantiation Parameters and the possible values/meanings
  // -- See micro-arch spec for more details
  //

  // The allowed parameters are as follows:
  // - ENA_ID48B: set to select how ID is defined. Note BCR/DCR separate
  // -  constants from ID_48B, ID_BCR, ID_DCR as needed
  // - ID_AS_REGS: controls which are SW controlled registers vs nets or constants
  // - ENA_SADDR: selects if there is an i2c static address and how provided if so
  //    constant from SADDR_P as needed
  // - ENA_CCC_HANDLING: selects which CCCs are handled by the block other than
  //                     the builtin ENTDAA, SETDASA, RSTDAA, SETNEWDA
  //    constants from MAX_DS_WR, MAX_DS_RD, MX_DS_RDTURN
  // - ENA_IBI_MR_HJ: enables event generation for interrupts, mastership/p2p, 
  //                     and hot join, as well as timing rule for Bus Avail match
  // - SEL_BUS_IF: related to use of bus and memory mapped registers
  // - PIN_MODEL: controls how pins are handled
  // - FIFO_TYPE: controls use of FIFO, whether internal or external
  // - EXT_FIFO: controls external FIFO type
  // - RSTACT: related to SlaveReset
  //
    // ENA_ID48B
  
  
  
  
    // ID_AS_REGS
  
  
  
  
  
  
  
  
  
  
    // ENA_SADDR
  
  
  
  
    // ENA_CCC_HANDLING
  
  
  
  
  
  
    // ENA_IBI_MR_HJ
  
  
  
  
  
  
  
    // ERROR_HANDLING
  
  
    // Time control enable bits
  
  
  
  
  
    // Time control SYNC values (mapped from commands)
  
  
  
  
  
  
  
    // HDR enable
  
  
    // SEL_BUS_IF
  
  
  
  
  
    // DMA_TYPE (when selected by SBIF_DMA_b)
  
  
  
  
    // PIN_MODEL
  
  
  
    // FIFO_TYPE
  
  
  
  
    // EXT_FIFO
  
  
    // RSTACT
  
  
  
  
  
  
    // ENA_MAPPED
    // Note that mapped used for i2c as well as extra (virtual) DA and SA
    //      Some i2c features are a sep param, but only if MAP enabled
  
  
  
  
  

  
  //
  // CCC related values
  //

   // params for CCC detection flags
  
  
  
  
  
  
  
  
  
  
  
  
          // 5 to 7 for DASA
  
  
  

  // params for CCCs
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  // now 1.1 version CCCs
  
  
  


  //
  // Interrrupt states (status)
  //
  
  
  
  
  
  
  
  
  
  
  
  

  //
  // Slave Reset action request value (what action to take)
  //
  
  
  
  

  //
  // Special observer for DFT - optional
  //
  

     






module i3c_sdr_slave_engine #(
    // params: ID and CCC ones not here since handled in separate module
    parameter         ENA_SADDR =  2'd0,    // none, const, reg/net
    parameter          SADDR_P  = 0,    // 7 bits as 7:1
    parameter         ENA_MAPPED= 5'd0, // if extra DAs/SAs allowed, and related
    parameter          MAP_CNT  = 4'd0, // number of extra DAs/SAs allowed
    parameter          MAP_I2CID= 24'd0,// !=0 if I2C extended with DevID
    parameter         ENA_IBI_MR_HJ = 0,// 0 if no events, else events as mask
    parameter         ENA_HDR   = 0,    // HDR-DDR modes support (only DDR here)
    parameter         PIN_MODEL =  0, // combinatorial pin use
    parameter         priv_sz   = PIN_MODEL[1] ? (ENA_HDR[0] ? 3 : 1) : 0,// wider if ext-pad+ddr
    parameter         ENG_DBG_MX=0      // observer
  )
  (
  // 1st Clock and Reset and hold
  input               clk_SCL_n,        // SCL falling edge: start and launching data
  input               clk_SCL,          // SCL rising edge: sample data and read T bit
  input               clk_SDA,          // SDA rising edge: STOP when SCL=1
  input               RSTn,             // master Reset
  input               SDR_hold,         // SDR hold when in HDR (see DDR nets below)
  input               cf_SdrOK,         // allows I3C SDR else i2c only
  // now SDA and SCL as pin inputs as well as outputs - note common mux for HDR
  // SDA is both push-pull and open-drain. 
  //   Open drain low:  oena=1, out=0
  //   Open drain high: oena=0, out=x -- is high-z
  input               pin_SCL_in,       // SCL as pin in: for IDLE detect and START/STOP
    // output         pin_SCL_out,      // SCL as output - not used here
    // output         pin_SCL_oena,     // SCL output selector - not used here
  input               pin_SDA_in,       // SDA as pin in: widely used
    // NOTE: the next 3 are different depending on PIN_MODEL. They are used as:
    //       == PINM_COMB: pin_SDA_out/oena connected to pads using shortest paths
    //       == PINM_REG:  pin_SDA_out/oena connected to pads with more leeway on paths
    //       == PINM_EXT_REG: pin_SDA_out/oena/oena_rise feed flops close to pads using
    //                        module i3c_pinm_reg_ext()
  output              pin_SDA_out,      // SDA as output: for read, ACK, and Arb
  output              pin_SDA_oena,     // SDA output selector: =1 for output
  output  [priv_sz:0] pin_SDA_oena_rise,// special on SCL rising if EXT_REG, else 0
  output              i2c_spike_ok,     // Is 1 to allow i2c spike filter
  // now some signals that indicate states to allow upper layer or system
  // act on them or interrupt. Sequence is int_start_seen, then possibly one of the
  // matching addresses (and fb_data_done), then data and possibly state_in_CCC, and
  // finally in_STOP.
  output              int_start_seen,   // 1 cycle pulse when START or repeated START seen
  output              int_start_err,    // held when SDA=1 after STOP - cleared on START
  output              int_da_matched,   // 1 cycle pulse if matched our dynamic address
  output              int_sa_matched,   // 1 cycle pulse if matched our static address (if any)
  output              int_7e_matched,   // 1 cycle pulse if matched i3c broadcast address
  output              int_in_STOP,      // STOP state (held until START)
  output              int_event_sent,   // 1 cycle pulse when event (event_pending) out
  output              int_event_ack,    // 1 cycle pulse with int_event_sent if ACK vs. NACK
  output              int_in_evproc,    // 1 if processing event now (so no cancel)
  // next is a temporary config
  input               cf_SlvNack,       // want us to NACK matches for now
  // now common buffers for to-bus. We use clk_SCL_n for the operational
  // clock for ack. The valid model 1 if data waiting, else 0
  // ack will allow it to advance (can read using SCL). 
  // Note exception is for DAA and CCC handled locally
  input               tb_data_valid,    // b1 if full, b0 if empty
  input         [7:0] tb_datab,         // input byte from system if valid==1
  input               tb_end,           // indicates no more data
  output              tb_datab_ack,     // ack for 1 SCL_n clock when consumed
  output              tb_urun_nack,     // pulse for 1 SCL_n if NACKed due to no data
  output              tb_urun,          // pulse for 1 SCL_n if no data and not end
  output              tb_term,          // terminated by master on read
  // now common buffers for from-bus. Layer above this uses clk_SCL for the operational 
  // clock to indicate buffer to fill. The use model is 01 (build lower), 10 (build upper), 
  // or 0 (is still full so if data in, will orun). 
  // Done pulsed by this block on clk_SCL_n. The data_use should not change until done.
  // Note that exception is for locally handled CCC
  input               fb_data_use,      // b1 if OK, b0 if full
  output        [7:0] fb_datab,         // output byte to system if done=1 and use=1
  output              fb_datab_done,    // done pulsed for one clk_SCL_n when written
  output              fb_datab_err,     // set with done if i3c parity error
  output              fb_orun,          // pulse for one clk_SCL_n if over-run (full when data in)
  output              fb_s0s1_err,      // S0 or S1 lockup error
  input               fb_s0s1_held,     // held error
  input               rd_abort,         // abort read due to clock stall (>60us)
  // now miscellaneous things related to CCCs, addresses, events
    // states used by upper layers for context
  output              state_accslv,     // 1 if read/write for slave
  output              state_rd,         // if we are writing output (read or IBI data)
  output              state_wr,         // if we are being written to (SDR write)
  output              in_daa_mode,      // state in ent daa
  output              state_dir,        // was_read from header on
    // CCC related for upper layers and local handling alike
  output        [7:0] state_in_CCC,     // defines detected CCC states/modes: see CF_xxx
  output        [1:0] next_ccc,         // preps for next broadcast or direct
  output        [1:0] next_hdr,         // will be HDR: [0]=hdr, [1]=ddr specifically
  output              in_ddr,           // in DDR specifically (note timing)
  input               ccc_handling,     // will be locally handled CCC
  input               ccc_handled,      // is locally handled CCC
  input               int_ccc_handled,  // is targeted to us
  input               ccc_uh_mask,      // mask for unhandled
  input               ccc_get,          // is a GET type directed
  output              ccc_is_read,      // if a read
  output              ccc_real_START,   // real START to cancel RSTACT
  output              int_ccc,          // CCC not handled by us, so app has to (or can NACK)
  output              done_rdy,         // for CCC handling (a cycle early since reg)
  output        [7:0] done_buff,        // this is used with CCC for the data from Master
  input               ccc_tb_now,       // have data to emit
  input               ccc_tb_data,      // actual data bit (same clock domain)
  input               ccc_tb_continue,  // 1 if continue, else is done (on 9th)
  output              is_9th,           // 1 for clock of 9th bit - no CCC data
  input               opt_s0ignore,     // ignore S0 error on request
    // Now sub-state of CCC for Dynamic Address Assignment "mode"
  output              daa_active,       // indicates when processing DAA; goes low if lost
  input               daa_inp_drv,      // if 1, then output daa_inp_bit
  input               daa_inp_bit,      // bit to output if in DAA and daa_inp_drv
  input               daa_done,         // pulse a cycle on ACK/NACK if we won (even if parity bad)
  input               daa_acknack,      // in ACK/NACK of ENTDAA new DA
    // addresses we can match - I3C DA and i2c SA (if used). Optional list
  input         [7:0] dyn_addr,         // dynamic address to match in [7:1] if [0]=1
  input         [7:0] opt_static_addr,  // static address to match in [7:1] if [0]=1 (SADDR controls)
  input [(MAP_CNT*10)-1:0] SetMappedDASA,// mapped SAs/DAs if supported
  input         [2:0] SetSA10b,         // if [1] is SA 10-bit
  output       [11:0] opt_match_idx,    // last 3 matched if used
  input               cf_vgpio,         // if CCC triggered VGPIO vs. sep DA
  output              vgpio_done,       // done pulse if VGPIO enabled
  output        [7:0] vgpio_byte,       // byte value if "
  input         [1:0] ccc_vgpio_done,   // when VGPIO is a CCC
  input         [7:0] ccc_vgpio,
  input               hdr_newcmd,       // is pulsed by HDR if CMD to go out (for MMR)
  output        [7:0] hdr_cmd,          // cmd if newcmd=1
    // i2c extended stuff if used
  input         [2:0] i2c_dev_rev,      // revision if DeviceID is used (OR into const)
  output              i2c_sw_rst,       // SW reset
    // nets related to events  
  input         [2:0] event_pending,    // !=0 if event (MR=3, IBI=2, HJ=1) pending. [2] is flag
  input               force_sda,        // force SDA as long as SCL=1. This looked at time
  input               ibi_has_byte,     // 1 if extra byte attached to IBI
  input         [7:0] opt_ibi_byte,     // byte to send if we have one
  input               ibi_timec,        // 1 if IBI is with Time control (with has_byte)
  output        [1:0] ibi_timec_marks,  // SC1 and SC2 stop - pulses
  output        [4:0] ibi_out_sel,      // selects source and index for IBI bytes
  input         [2:0] ibi_extdata,      // if switch to read after IBI byte
  input         [3:0] ibi_mapidx,       // mapped index for IBI if mapping used
  output              ibi_in_extdata,   // if reading from ext data
    // now nets relates to HDR-DDR if in use (see SDR_hold)
    // This engine only handles the bytes. Preamble and parity are
    // handled by DDR block. So, this uses 3 sets of flops: bit_cnt
    // fb_datab_r, and ccc_byte. All handshakes are handled outside
  input         [1:0] ddr_fb_ctrl,      // 0=collect, 1=push LSB, 2=push MSB
  input               ddr_fb_MSB,       // from-bus DDR 1st byte
  input               ddr_fb_LSB,       // from-bus DDR 2nd byte
  input         [1:0] ddr_tb_ctrl,      // 0=off, 1=use tbuff, 2=use our bits, 3=copy MSB to extra
  input         [1:0] ddr_tb_bits,      // pre-amble & parity for rising and falling
  input         [2:0] ddr_tb_oe,        // handles OE for preamble
  output        [2:0] obit_cnt,         // bit counter from engine
  output        [7:0] byte_1st,         // 1st byte 
  output        [7:0] byte_2nd,         // 2nd byte 
  output[ENG_DBG_MX:0]eng_debug_observ, // optional debug observer
  input               scan_no_rst       // prevents layered reset
  );

  // -------------------------
  // Table of Contents
  // 1. Above: ports
  // 2. Special clock domain for SDA to handle STOP
  //    -- Note that this clock domains collapses to SCL in scan
  // 3. Special clock domain for SCL (vs. SCL_n) 
  //    -- Note that this clock domains collapses to SCL in scan
  // 4. Define registers and wires for normal SCL_n clock domain
  // 5. Top level state machine for i2c and SDR
  //    -- Note special override from registered START detect
  // 6. Flops and controls in service to state machine
  //    -- includes state and data handling
  // 7. Tracking mechanism related to CCCs. 
  //    -- Actual CCC handling in a separate block, but this
  //       tracks CCC state (end depends on type of CCC) and
  //       the two CCC modes (DAA and HDR) as well as
  //       doing some decoding.
  // 8. Pad output control for SDA
  //    -- SDR and i2c only affect SDA, not SCL
  //    -- Both OD and PP (std, as used by i3c)
 
  //
  // 2. Special clock domain for SDA to handle STOP
  //    -- Note that this clock domains collapses to SCL in scan
  //
  // STOP is triggered by SDA going High when SCL is High and we are not
  //      being suppressed. It is reset by START detect. 
  // Note: START and repeated START is handled with falling edge of SCL
  reg                 STOP_r;
  wire                rst_STOP_n;
  wire                any_START_d;      // START and repeated START (SCL_n)
  wire                real_START_d;     // START after STOP (SCL_n)
  reg                 init_stop;        // 1st time we come up is init stop (SCL)

  // STOP is SDA going High when SCL is High
  // Note ugly test of pin_SCL_in for High. This is because they can
  // do START, STOP such that STOP is just after rising after START, so
  // reset would not clear (until rising of SCL). We cannot use registered
  // on SCL rising, because START that follows would not have cleared
  // rising. 
  // Note: STOP resets on START or likely START
  assign rst_STOP_n  = RSTn & (~any_START_d | pin_SCL_in | scan_no_rst); 
   reg observe_STOP_rst ; \
        always @ (posedge  clk_SCL) observe_STOP_rst <=  ~any_START_d ^ pin_SCL_in ^ ~STOP_r  ^ ~rd_abort;  // optional DFT observer

  always @ (posedge clk_SDA or negedge rst_STOP_n)
    if (!rst_STOP_n)
      STOP_r      <= 1'b0;              // cannot be 1 due to any_START
    else if (pin_SCL_in & ~SDR_hold)    // SDA rise when SCL=high
      STOP_r      <= 1'b1;              // stopped as a "state"
  assign int_in_STOP = STOP_r | init_stop; // outer system needs to know

  //
  // 3. Special clock domain for SCL (vs. SCL_n) 
  //    -- Note that this clock domains collapses to SCL in scan
  //
  reg            SDA_r;                 // registered on SCL rising, used on falling
  // we register SDA on each SCL rising. Need for M->S addr, data, IBI ACK, but also Sr detect
  always @ (posedge clk_SCL or negedge RSTn)
    if (!RSTn)
      SDA_r <= 1'b0;
    else if (fb_s0s1_held)
      SDA_r <= 1'b1;                    // in prep for START
    else if ((~SDR_hold|next_hdr[0]) | ddr_fb_MSB | ddr_fb_LSB)
      SDA_r <= pin_SDA_in;              // simply register SDA on each
  // FUTURE: need to see if better way to detect 1st time coming up STOPped
  always @ (posedge clk_SCL or negedge RSTn)
    if (!RSTn)
      init_stop <= 1'b1;
    else if (~SDR_hold)
      init_stop <= 1'b0;

  //
  // 4. Define registers and wires for normal SCL_n clock domain
  //
  reg            start_possible;        // SDA low on SCL falling
  reg            start_any;             // full START or repeated START
  reg            start_invalid;         // detected post STOP non-start
  reg            start_r;               // explicitly START, not repeated
  wire           addr_match;            // address is one of ours
  wire           map_da_match, map_sa_match; // special in case mapped SA/DA
  wire           in_sa2byt;             // matched the 7'b11110_bb or 0x0 part
  wire           matched_7e;            // i3c broadcast addr match
  wire           s0_error, s0_error2;
  wire           matched_p2p;           // i3c point to point fixed addr=7'h1
  reg            was_read;              // registered with RnW
  reg      [2:0] bit_cnt;               // bit counter for byte build
  reg      [7:0] fb_datab_r;            // FB 1 or 2 bytes: selected by upper layer
  reg      [7:0] ccc_byte;              // used for CCCs so not failed if buff busy
  wire           use_ccc_byte;          // indicates if we this or datab
  reg            use_ccc_r;
  wire     [7:0] work_datab;            // current or just last
  reg            datab_done;            // 1 clock pulse for done with FB
  reg            datab_orun;            // upper layer was not ready when we needed fb
  wire           parity;                // parity on data
  wire           rdata_bit;             // 1 bit of data for S->M
  wire     [1:0] rddr_bits;             // used if in DDR read mode
  reg            daa_lost;              // 1 if we lost the arb on DAA
  reg            event_won;             // 1 if event arb won (or have not lost yet)
  wire           event_losing;          // 1 if losing arb
  wire           event_drive;           // indicates ARB output on START header
  wire     [7:0] event_hdr;             // address and RnW to use
  wire           event_hdr_b;           // current bit
  wire           force_START;           // special for IBI when no clock
  reg      [3:0] extdata_pos;           // index into time-control and extdata
  reg            ext_end;               // cycle delayed end
  reg            ibi_rd;                // used for tracking IBI data
  wire           ack_ok;                // allows us to NACK if cannot accept 
  reg            s4_err;
  wire           map_spc;               // ACK for i2c special
  wire           vgpio_spc;             // ACK for VGPIO
  wire           vgpio_msg;             // if matched VGPIO addr
  wire           vgpio_ccc;             // if a VGPIO CCC, so handled
  wire           dev_spc;               // ACK for i2c DEVID special
  wire           spc_data_valid;        // used for i2c device ID (const) if used, else 0
  wire           spc_data_end;          // "
  wire     [7:0] spc_data;              // 8 bits of device ID if used
  reg      [7:0] in_ccc;                // tracks CCC as mode/state only
  reg      [1:0] ccc_supp;              // for signaling CCC to slave
  reg            ccc_uh_r;              // registered once accepted
  wire           ccc_valid_dir;         // directed CCC direction is valid
  wire           next_int_ccc, next_bc_ccc, next_dc_ccc;
  reg            pass_ccc;
  wire           is_i3c = dyn_addr[0] & cf_SdrOK;  // just shorthand for i2c vs. i3c
  wire           i3c_protocol = cf_SdrOK & (is_i3c | (|in_ccc));
  reg            in_i3c_msg;            // registers after addr if 7E or i3c_protocol
  wire           pend_IBI = event_pending[1:0]==2'h1; // interrupt
  wire           pend_MR  = event_pending[1:0]==2'h2; // master request 
  wire           pend_HJ  = event_pending[1:0]==2'h3; // hot join (no DA)
  wire           i3c_end  = (ccc_handled ? ~ccc_tb_continue : (tb_end|~tb_data_valid));
  reg            end_flag;              // to-bus end marker (so we remember)
  reg            is_term;               // is terminated by Master (for i3c case)
  wire           will_term;             // will register as abort by M
  wire           i2c_nack;
  wire           end_read;
  wire           ibi_exttx;             // extended IBI data and TXFIFO not empty
  wire           SDA_oe, SDA_out;       // pad drive
  wire           r_SDA_oena;            // only used for reg pad ctrl

  //
  // 5. Top level state machine for i2c and SDR
  //    -- Note special override from registered START detect
  //
  reg      [3:0] state_base;
  wire     [3:0] state;
  reg      [3:0] next_state;
   
    // 1st the cases from master (except ARB OD for IBI/MR/HJ)
    // Do not change the numbering without understanding relationship of bits[3:2]
    // and mapping of A7_A0 for event and non-event
  localparam ST_WAIT_SrP      = 4'b0000; // not us, so we ignore until (re)START/STOP
  localparam ST_WRITE         = 4'b0001; // M->S
  localparam ST_A7_A0_RnW     = 4'b0011; // Address header from (re)START
    // now the ones where we drive the bus (bits [3:2] != 0)
  localparam ST_IBI_BYTE      = 4'b0100; // IBI data byte if used
  localparam ST_IBI9TH        = 4'b0101; // IBI data 9th T bit - 0 unless EXTDATA
  localparam ST_MACK          = 4'b0110; // IBI/MR/HJ M->S ACK of our addr and handoff
  localparam ST_ACK_NACK      = 4'b1000; // We ACK (or not) address header
  localparam ST_READ          = 4'b1001; // S->M and Peer-to-Peer
  localparam ST_EV_A7_A0_RnW  = 4'b1011; // Event + Address header from (re)START
  localparam ST_R9TH          = 4'b1100; // i2c: M->S ACK, i3c: T (end or allow term)
  localparam ST_W9TH          = 4'b1101; // i2c: S->M ACK (so OD), i3c: input T (parity) 
  localparam ST_DAA           = 4'b1110; // ENTDAA mode - OD arb and not 9th bit


  //
  // Allow export of debug observer data
  // -- default stub assigns 0, but this allows adding nets
  //
  



   assign eng_debug_observ = {1'b0}; // no observer output by default
  

  //


  //
  // Main state machine: this operates the top level states for SDR
  // -- Sub FSMs handle counting out bits and the like
  // -- State machine operates off falling edge of SCL, SDA registered 
  //    on previous rising. So, other than repeated-start and stop,
  //    this simplifies everything.
  // -- Special case is repeated START, which is psuedo combinatorial
  //
  wire state_rst_n = RSTn & (scan_no_rst | ~rd_abort);
  always @ (posedge clk_SCL_n or negedge state_rst_n) 
    if (!state_rst_n)
      state_base <= ST_WAIT_SrP;
    else if (SDR_hold)
      state_base <= ST_WAIT_SrP;             // note that SCL going low without SDA low 1st is junk
    else
      state_base <= next_state;
  // next forces state into A7 if START happened just before falling edge. It 
  // also handles STOP in the middle of a message followed by SCL low
  assign state        = s4_err ? ST_WAIT_SrP :
                        any_START_d ? (event_drive ? ST_EV_A7_A0_RnW : ST_A7_A0_RnW) : 
                                      STOP_r ? ST_WAIT_SrP : state_base;
  assign state_accslv = |state & ~int_in_STOP; // is actively involved
  assign state_rd     = (^state[3:2] | (state==ST_R9TH)) & |state[2:0] & ~int_in_STOP; // we are writing 
  assign state_wr     = ((state==ST_WRITE)|(state==ST_W9TH)) & ~int_in_STOP; // M writing data
  assign is_9th       = state == ST_R9TH;
  assign state_dir    = was_read;

  // State
  always @ ( * ) 
    case (state)
    // IDLE:                            // idle not possible as STOP is not a clocked state
    ST_A7_A0_RnW,                       // Address header - runs until RnW
      ST_EV_A7_A0_RnW:                  // same but with event output
      // 7 bits of address plus RnW makes 8. This could be Master or us driving (arb for IBI)
      // event_won is if we were trying IBI/MR/HJ and won the arb
      // addr_match is if matches 7'h7E or our SA/DA (depending if DA) 
      next_state = |bit_cnt ? ((event_won&~event_losing) ? ST_EV_A7_A0_RnW : ST_A7_A0_RnW) : 
                   (event_won&~event_losing) ? ST_MACK : addr_match ? ST_ACK_NACK : 
                    ST_WAIT_SrP;        // if not for us, wait for reSTART/STOP

    ST_ACK_NACK:                        // we ACK if to us (our address or 7'h7E)
      // ack_ok tells us if we can accept the read or write; if not, we NACK (do nothing)
      // read is read from us, write is write from master. Daa is ENTDAA mode for
      // ID processing and happens in_daa mode when 0x7E/R is seen; as well as S4 error
      next_state = (in_daa_mode & was_read) ? ST_DAA : ~ack_ok ? ST_WAIT_SrP :
                   was_read ? ((tb_data_valid | ccc_handled) ? ST_READ : ST_WAIT_SrP) : 
                   ST_WRITE; 

    ST_READ:                            // read 8 bits from us then ACK or T  
                                        // is_term from i3c read abort
      next_state = (is_term | rd_abort) ? ST_WAIT_SrP : |bit_cnt ? ST_READ : ST_R9TH;

    ST_R9TH:                            // 9th bit is M ACK/NACK if i2c, T if i3c
      // note: i3c terminate in ST_READ. i2c NACK here
      next_state = (end_read | i2c_nack | rd_abort) ? ST_WAIT_SrP : ST_READ; 
   
    ST_WRITE:                           // write 8 bits to us and then ACK or T
      // note: this may be "fake bit D7" to change SDA for repeated START or STOP
      next_state = |bit_cnt ? ST_WRITE : ST_W9TH;

    ST_W9TH:                            // 9th bit is our ACK/NACK if i2c, T parity if i3c
      next_state = (in_i3c_msg&(parity!=SDA_r)) ? ST_WAIT_SrP : ST_WRITE; 

    ST_DAA:
      // we arbitrate with our ID if we are not assigned yet. If we win, we get a DA
      // and then ACK. Done marked if we won even if we NACK (bad parity)
      next_state = (daa_lost | daa_done) ? ST_WAIT_SrP : ST_DAA;

    ST_MACK:                            // if IBI or MR or HotJoin, this is Master ACK/NACKing us
      next_state = SDA_r ? ST_WAIT_SrP :// not accepted IBI/MR/HJ 
                   (pend_IBI & ibi_has_byte) ? ST_IBI_BYTE : // IBI byte follows
                   ST_WAIT_SrP;         // done our job, in Master's hands now
                   

    ST_IBI_BYTE:                        // emit byte for IBI
      // this could be in ST_READ
      next_state = |bit_cnt ? ST_IBI_BYTE : ST_IBI9TH;

    ST_IBI9TH:
      next_state = (~|extdata_pos | ext_end) ? 
                        (ibi_exttx ? ST_READ : ST_WAIT_SrP) : // last byte
                                 ST_IBI_BYTE;
 
    ST_WAIT_SrP:
      next_state = ST_WAIT_SrP;         // this will stay here until START or repeated START

    default:
      next_state = ST_WAIT_SrP;         // this will stay here until START or repeated START
    endcase

  // register START and repeated START - affects state machine and A7 actions
  // any_START_d clears is_STOP. 
  // Note: we break into separate flops so metastable SDA/SCL does not
  // propagate into logic; we use post reg logic with a very shallow
  // cone, so valid just after Qs
  always @ (posedge clk_SCL_n or negedge RSTn)  
    if (!RSTn) begin
      start_possible <= 1'b0;
      start_any      <= 1'b0;
      start_invalid  <= 1'b0;
      start_r        <= 1'b0;
    end else if (~SDR_hold) begin       // not when in HDR
      start_possible <= ~pin_SDA_in;    // needed for repeated START and START
      start_any      <= (int_in_STOP | SDA_r); // needed with SDA=0
      start_invalid  <= pin_SDA_in & int_in_STOP; // not a valid combo. 
      // start_r is not used unless any_START_d is true, so no risk
      if (int_in_STOP)
        start_r      <= 1'b1;           // special from START (from Bus Available)
      else if (~pin_SDA_in & SDA_r)
        start_r      <= 1'b0;           // repeated START
    end 
  assign any_START_d    = start_possible & start_any;
  assign real_START_d   = any_START_d & start_r;
  assign ccc_real_START = real_START_d;
  assign int_start_seen = any_START_d;  // outer system needs to know
  assign int_start_err  = start_invalid;// held until START


  //
  // 6. Flops and controls in service to state machine
  //    -- includes state and data handling
  //
  // read in RnW bit of header
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      was_read <= 1'b0;
    else if ((state[2:0] == ST_A7_A0_RnW[2:0]) & (bit_cnt == 3'h0))
      was_read <= SDA_r;

  // the bit counter is used to serialize and deserialize the data
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      bit_cnt   <= 3'h7;
    else if (~SDR_hold) begin
      // NOTE: If SDA is too close (meta-stable) to SCL on START, this could end up with
      //       junk bit_cnt since not. all bits would change the same. That is OK. As it
      //       just means no match on address.
      if (~pin_SDA_in & (int_in_STOP | SDA_r))
        bit_cnt <= 3'h7;
      else if ((state[2:0] == ST_A7_A0_RnW[2:0]) | (state == ST_WRITE) | 
               (state == ST_IBI_BYTE) | (state == ST_READ) |
               ((PIN_MODEL!= 0) & (state == ST_ACK_NACK))) // to get the 7
        if ((state != next_state) & (next_state != ST_A7_A0_RnW))
          bit_cnt <= 3'h7;
        else
          bit_cnt <= bit_cnt - 3'h1;
    end else if (ENA_HDR[ 0]) begin
      // if DDR enabled, we count
      if (ddr_fb_MSB | ddr_fb_LSB)
        bit_cnt <= bit_cnt - 3'h1;
      else if (ddr_tb_ctrl[0])
        bit_cnt <= bit_cnt - 3'h1;
      else
        bit_cnt <= 3'h7;
    end

  // TO-BUS: when it is Read (send to Master), we need to use tb_data
  //         buffer. Need to ack on 9th and use tb_end to
  //         signal end if i3c and drive err if i3c and not NACKed
  // Note:   on i3c read high-Z, the master can terminate. This would be a
  //         a repeated start, so we would abandon anyway.
    // rdata_bit just uses bit from correct buffer, or High if none
  assign rdata_bit    = ccc_handled ? ccc_tb_data :
                        spc_data_valid? spc_data[bit_cnt] : // only if i2c DevID
                        tb_data_valid ? tb_datab[bit_cnt] : 
                        1'b1;           // undriven if i2c and just 1 if i3c
  generate if (ENA_HDR[ 0]) 
    assign rddr_bits  = ddr_tb_bits[1:0];
  else
    assign rddr_bits  = 2'b00;
  endgenerate

  // next 3 signals back up flag on ACK and read of data
  assign tb_datab_ack = ((state==ST_R9TH)| &ddr_tb_ctrl) & 
                        (~spc_data_valid & tb_data_valid) & ~ccc_handled;// signal done with buffer
  wire   read_urun    = (state==ST_READ) & ~tb_data_valid & ~ccc_handled & ~is_term;
  assign tb_urun      = read_urun & ~ibi_rd;
  assign tb_urun_nack = (state==ST_ACK_NACK) & was_read & ~tb_data_valid &
                        ~ccc_handled & ~matched_7e;// DAA is special
  // Term detects when i3c master aborts while SDA is floating (SCL high)
  // and when i2c master NACKs (SCL high). 
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      is_term <= 1'b0;
    else if ((state==ST_R9TH) & ~end_read)
      is_term <= will_term;
    else if (is_term)
      is_term <= 1'b0;
  assign tb_term      = is_term; // under run is more important
  assign i2c_nack     = ~in_i3c_msg & SDA_r;
  assign will_term    = in_i3c_msg ? ~pin_SDA_in : SDA_r;
  assign end_read     = ccc_handled ? ~ccc_tb_continue : end_flag;
  // end flag picks up the end marker with the data (or lack of).
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      end_flag <= 1'b1;
    else if (tb_datab_ack | (state==ST_ACK_NACK) | (next_state==ST_R9TH))
      end_flag <= tb_end | ~tb_data_valid; 

  // FROM-BUS: we serialize input from master, including address
  // Note that we use ccc_byte when in a handled CCC and when not in a CCC
  // and is header. This prevents running out of buffer space
  wire not_app_ccc = (ccc_handled | ~in_ccc[ 2]);
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      use_ccc_r   <= 1'b1;              // use by default
    else if (~SDR_hold)
      if ((state[2:0]==ST_A7_A0_RnW[2:0]) & not_app_ccc)
        use_ccc_r <= 1'b1;              // header, unless in cmd not handled
      else if ((state==ST_ACK_NACK) & ~ccc_handled & ~matched_7e)
        use_ccc_r <= 1'b0;              // only if handled CCC after this
      else if ((state==ST_ACK_NACK) & ~ccc_handled & matched_7e)
        use_ccc_r <= 1'b1;              // we use ccc_byte for CCC code
      else if (((state==ST_W9TH) | (state==ST_R9TH)) & ~ccc_handled)
        use_ccc_r <= 1'b0;              // has CCC byte if needed
  assign use_ccc_byte = use_ccc_r | real_START_d | (any_START_d & not_app_ccc);

  // from-bus data buffer is copied into ping-pong or FIFO by upper layers
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      fb_datab_r            <= 8'd0;
    else if (next_bc_ccc | next_dc_ccc) begin // provide CCC
      fb_datab_r[7:0]       <= ccc_byte;
    end else if (~SDR_hold & ~use_ccc_byte) begin
      if (((state[2:0] == ST_A7_A0_RnW[2:0]) | (state == ST_WRITE)))
        fb_datab_r[bit_cnt] <= SDA_r;
    end else if (ddr_fb_LSB & ENA_HDR[ 0]) begin
      // 2nd byte: used if DDR is enabled
      fb_datab_r[{bit_cnt[1:0],1'b1}] <= SDA_r;
      fb_datab_r[{bit_cnt[1:0],1'b0}] <= pin_SDA_in;
    end else if ((ddr_fb_ctrl==2'b01) & ENA_HDR[ 0])
      fb_datab_r[7:0]       <= ccc_byte; // move MSB down to push to FIFO
  assign hdr_cmd = fb_datab;            // what would be pushed to FIFO

  // ccc byte is 8b for specialized uses
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      ccc_byte            <= 8'd0;
    else if (~SDR_hold & use_ccc_byte) begin
      if (((state[2:0] == ST_A7_A0_RnW[2:0]) | (state == ST_WRITE)))
        ccc_byte[bit_cnt] <= SDA_r;
    end else if (ddr_fb_MSB & ENA_HDR[ 0]) begin
      // 1st byte: used if DDR is enabled
      ccc_byte[{bit_cnt[1:0],1'b1}] <= SDA_r;
      ccc_byte[{bit_cnt[1:0],1'b0}] <= pin_SDA_in;
    end
  assign work_datab= use_ccc_byte ? ccc_byte : fb_datab_r[7:0];
  assign fb_datab  = fb_datab_r[7:0];
    // buffer we just finished (used by CCC handler)
  assign done_buff = (state==ST_WRITE) ? {work_datab[7:1],SDA_r} : work_datab; 
    // next three only used if DDR is enabled
  assign obit_cnt  = bit_cnt;
  assign byte_2nd  = ccc_byte;
  assign byte_1st  = fb_datab_r;

  // on data inbound completion, mark done and also compute parity (i3c)
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      datab_done  <= 1'b0;              // i3c: set on 9th, i2c: set on 8th
    else if (~|bit_cnt & (state == ST_WRITE) & (~in_ccc[0] | ccc_uh_r)) 
      datab_done  <= 1'b1;              // data done, for i3c parity can be error
    else 
      datab_done  <= 1'b0;              // 1 cycle pulse whether i2c or i3c
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) 
      datab_orun  <= 1'b0;              // tracks if no buffer when we need it
    else if (~STOP_r & (&bit_cnt[2:0]) & ~fb_data_use &
                 ((state == ST_WRITE) | (state[2:0] == ST_A7_A0_RnW[2:0])))
      datab_orun  <= 1'b1;              // upper layer not ready for us
    else if (|ddr_fb_ctrl & ~fb_data_use)
      datab_orun  <= 1'b1;              // upper layer not ready for us
    else 
      datab_orun  <= 1'b0;              // overun is 1 cycle pulse
  assign done_rdy      = (state == ST_W9TH) & in_i3c_msg;
    // the following is combinatorial for the ccc stuff because we
    // have to get fb_data_use switched before real data in next cycle
    // FUTURE: may want to register some of these below to avoid
    //         potentially slow cone
  assign fb_datab_done = ~ccc_handled & 
                         ((~use_ccc_byte & datab_done) | pass_ccc | (|ddr_fb_ctrl)) & 
                         ~in_sa2byt & ~dev_spc & // no store if 2nd byte of ext i2c
                         ~vgpio_ccc & ~vgpio_msg;  // special - registered up link
    // S2 error is not to be signaled until SDA_r is in (and paired to done)
  assign fb_datab_err  = (state==ST_W9TH) & in_i3c_msg & (parity!=SDA_r);
    // S0/S1 error (if not suppressed) are exported to lockup until Exit or timeout
  assign fb_s0s1_err   =  ~opt_s0ignore & (s0_error | s0_error2 | 
                            (fb_datab_err & (in_ccc[ 4:0]== 5'd1)));
    // next signals that we over-ran (needed it to be free and wasn't). Note
    // that we can overrun on address header, even though we do not pass up
  assign fb_orun   = datab_orun;
  assign parity    = ^work_datab ^ 1'b1; // odd parity

  // next one is held during message if I3C format
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      in_i3c_msg   <= 1'b0;
    else if ((state[2:0] == ST_A7_A0_RnW[2:0]) & (bit_cnt==3'd1))
      if (i3c_protocol | ({work_datab[7:2],SDA_r} == 7'h7E))
        in_i3c_msg <= 1'b1;
      else
        in_i3c_msg <= 1'b0;

  generate if (ENA_MAPPED[ 0]) begin : match_mapped
    genvar           i;
    localparam BASE = ENA_MAPPED[ 4] ? 2 : 1;
    wire [MAP_CNT:BASE] mat_sa;
    wire [MAP_CNT:BASE] mat_da;
    wire             mat_ext, gcall_addr;
    reg       [11:0] matched_map_r;
    wire       [3:0] found[1:15];
    wire             sa_10b_read, sa2_byte, sa2_skip;
      // below: 1st one is each address (and enable), 2nd one is SA/DA and NACK
    wire [(MAP_CNT*8)-1:0] map_regs_ad  = SetMappedDASA[(MAP_CNT*10)-1:(MAP_CNT*2)];
    wire [(MAP_CNT*2)-1:0] map_regs_det = SetMappedDASA[(MAP_CNT*2)-1:0];
    for (i = BASE; i <= MAP_CNT; i = i + 1) begin : mat_sada
      wire [7:0] loc = map_regs_ad[(((i-1)*8)+7) -: 8];
      wire [1:0] sa_nack = map_regs_det[(((i-1)*2)+1) -: 2];
       // we don't match if NACK or wrong type (DA vs. SA)
      assign mat_da[i] = loc[0] & ~|sa_nack & (loc[7:1]==work_datab[7:1]);
      assign mat_sa[i] = loc[0] & sa_nack[0] & ~sa_nack[1]  & (loc[7:1]==work_datab[7:1]);
      assign found[i]  = (mat_da[i] | mat_sa[i]) ? i : 4'd0; // only 1 will get set
    end

    if (ENA_MAPPED[ 4]) begin : vgpio_handling
      // Specialty VGPIO - one group only
      reg [7:0] vgpio_r;
      reg       vgpio_mat_r;
      reg       vgpio_done_r;
      wire [7:0]loc         = map_regs_ad[(((1-1)*8)+7) -: 8];
      wire      is_vgpio    = loc[0] & (loc[7:1]==work_datab[7:1]) &
                              ~cf_vgpio; // is CCC if =1
      assign vgpio_done     = vgpio_done_r;
      assign vgpio_byte     = vgpio_r;
      assign vgpio_spc      = is_i3c & ~cf_SlvNack & is_vgpio;
      assign vgpio_msg      = vgpio_mat_r; // in a GPIO message

      assign vgpio_ccc      = ccc_vgpio_done[1];
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          vgpio_mat_r <= 1'b0;
        else if (state==ST_ACK_NACK)
          vgpio_mat_r <= is_i3c & ~cf_SlvNack & is_vgpio;

      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn) begin
          vgpio_r      <= 8'd0;
          vgpio_done_r <= 1'b0;
        end else if (vgpio_mat_r & datab_done) begin
          vgpio_r      <= fb_datab;
          vgpio_done_r <= 1'b1;
        end else if (cf_vgpio & ccc_vgpio_done[0]) begin
          vgpio_r      <= ccc_vgpio;
          vgpio_done_r <= 1'b1;
        end else if (vgpio_done_r)
          vgpio_done_r <= 1'b0;

      assign found[1]       = 4'd0;
    end else begin
      assign vgpio_done     = 1'b0;
      assign vgpio_byte     = 8'd0;
      assign vgpio_spc      = 1'b0;
      assign vgpio_msg      = 1'b0;
      assign vgpio_ccc      = 1'b0;
    end

    if (ENA_MAPPED[ 1]) begin : extended_i2c
      // extended 10 bit address
      wire           mat_ext_2nd;
      reg            match_sa_10b;
      reg            match_ext_1st;
      reg            skip_part, check_2nd;
      // extended 10-bit only possible in map [1] if SA and extenstion is not 0
      // Some oddities in model. 1st wNe match 7'b111_10_xx where xx has to be
      // the 2 MSbs. Then next byte is remaining 8 bits. We have to ACK each.
      // So, reg match_ext_1st set if 1st part matched so we check 2nd, but only
      // if write. Then we reg in match_sa_10b until STOP. If match of 7'b11110xx
      // but read, then only match if we were valid last using maych_sa_10b
      wire [7:0] loc1    = map_regs_ad[(((1-1)*8)+7) -: 8];
      wire [1:0] sa_nack1= map_regs_det[(((1-1)*2)+1) -: 2];
      assign mat_ext     = loc1[0] & sa_nack1[0] & ~sa_nack1[1] & |SetSA10b &
                            (work_datab[7:3]==5'b11110) &
                            (work_datab[2:1]==SetSA10b[2:1]); // MSb match
        // next is for read after in same frame
      assign sa_10b_read = match_sa_10b & 
                           (work_datab[7:3]==5'b11110) &
                           (work_datab[2:1]==SetSA10b[2:1]) & work_datab[0];
      assign mat_ext_2nd = work_datab[7:0] == {SetSA10b[0],loc1[7:1]};
        // next signals a match and controls the ACK
      wire   is_2nd      = ((state==ST_W9TH) & match_ext_1st & mat_ext_2nd & check_2nd) | 
                           ((state==ST_ACK_NACK) & sa_10b_read);
        // note on below: this stops it storing into the FIFO
      assign sa2_skip    = skip_part | check_2nd;
      assign sa2_byte    = is_2nd | sa2_skip; // 2nd B match, read (so implicit), or skip
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          match_ext_1st  <= 1'b0;
        else if (is_i3c) begin
          if (match_ext_1st)
            match_ext_1st<= 1'b0;
        end else if ((state==ST_ACK_NACK) & ~is_i3c & ~cf_SlvNack)
          match_ext_1st  <= mat_ext & ~work_datab[0]; // is match 1st half (W)
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          match_sa_10b  <= 1'b0;
        else if (~is_i3c & match_ext_1st & (state==ST_W9TH) & mat_ext_2nd)
          match_sa_10b  <= 1'b1;
        else if (real_START_d)          // clear after STOP (so on real START)
          match_sa_10b  <= 1'b0;
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          check_2nd     <= 1'b0;
        else if ((state==ST_ACK_NACK) & ~is_i3c & ~cf_SlvNack & mat_ext & ~work_datab[0])
          check_2nd     <= 1'b1;        // check 2nd byte after match ext 1st
        else if (check_2nd & (state==ST_W9TH))
          check_2nd     <= 1'b0;        // cancel after that byte
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          skip_part     <= 1'b0;
        else if ((state==ST_W9TH) & check_2nd & ~mat_ext_2nd)
          skip_part     <= 1'b1;        // not our 10b, but matches 1st part
        else if ((state==ST_ACK_NACK) | is_i3c | cf_SlvNack)
          skip_part     <= 1'b0;        // cancel
    end else begin
      assign mat_ext     = 1'b0;
      assign sa2_byte    = 1'b0;
      assign sa2_skip    = 1'b0;
      assign sa_10b_read = 1'b0;
    end
    if (ENA_MAPPED[ 2]) begin : i2c_slvrst
      // I2C slave reset model
      reg            gen_call, SW_reset;
      wire           gcall;
      // note on below: this matches gen-call with Read, which is not tech correct
      assign gcall_addr = ~|work_datab[7:1];
      // we monitor for general-call and ACK
      wire isgcall = (state==ST_ACK_NACK) & ~is_i3c & ~cf_SlvNack &
                     gcall_addr & ~was_read; // general call
      assign gcall = isgcall | ((state==ST_W9TH) & gen_call);
        // reg the gen-call so we match next byte
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          gen_call  <= 1'b0;
        else if (is_i3c | (state==ST_W9TH)) begin
          // GEN call is always 1 byte after
          if (gen_call)
            gen_call<= 1'b0;
        end else if (state==ST_ACK_NACK)
          gen_call  <= isgcall;          // clears on next addr
      // Note: SW reset can be set, but is only cleared by new gen call
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          SW_reset <= 1'b0;
        else if ((state==ST_W9TH) & gen_call)
          SW_reset <= work_datab==8'h06; // SW reset is gen-call then 6
      assign i2c_sw_rst  = SW_reset; 
      assign in_sa2byt   = sa2_byte | gen_call;
    end else begin
      assign gcall_addr  = 1'b0;
      assign i2c_sw_rst  = 1'b0;        // unused;
      assign in_sa2byt   = sa2_byte;
    end

    for (i = MAP_CNT+1; i < 16; i = i + 1) begin : mat_fake
      assign found[i] = 4'd0;
    end
    // the next line is dumb but generate does not make it easy to build up a one 
    // hot value which is wider than a bit
    wire [3:0] this_match = found[1]|found[2]|found[3]|found[4]|found[5]|found[6]|found[7]|
                            found[8]|found[9]|found[10]|found[11]|found[12]|found[13]|found[14]|found[15];
    assign map_da_match  = |mat_da;
    assign map_sa_match  = |mat_sa | (in_sa2byt & ~sa2_skip);
    assign map_spc       = gcall_addr | mat_ext;
    assign opt_match_idx = matched_map_r;
    always @ (posedge clk_SCL_n or negedge RSTn) 
      if (!RSTn)
        matched_map_r <= 12'd0;
      else if (int_da_matched | (int_sa_matched & ~mat_ext))
        matched_map_r <= {matched_map_r[7:0],this_match[3:0]};
  end else begin
    assign map_da_match  = 1'b0;
    assign map_sa_match  = 1'b0;
    assign map_sa_idx    = 4'd0;
    assign map_spc       = 1'b0;
    assign in_sa2byt     = 1'b0;
    assign i2c_sw_rst    = 1'b0;        // unused;
    assign vgpio_done    = 1'b0;
    assign vgpio_byte    = 8'd0;
    assign vgpio_spc     = 1'b0;
    assign vgpio_msg     = 1'b0;
    assign vgpio_ccc     = 1'b0;
  end
  if (|MAP_I2CID) begin : i2c_dev_id
    reg           st_dev_id;
    reg     [1:0] dev_cnt;

    wire dev_id    = work_datab[7:1]==7'h7C; // DeviceID req
    wire isdevid   = ~is_i3c & ~cf_SlvNack & dev_id; // device ID 1st part
                      // test below is for data after 7C and so not in normal place
    wire ourid     =  opt_static_addr[0]&(work_datab[7:1]==opt_static_addr[7:1]);
    assign dev_spc = isdevid | (st_dev_id & (state==ST_W9TH) & ourid) |
                     (st_dev_id & (state==ST_R9TH) & |dev_cnt);
    // reg the steps (START,7C/W,our_addr,Sr,7C/R,<24b ID>,...
    wire dev_rst_n = RSTn & (~STOP_r | scan_no_rst);  // no observe since used elsewhere
    always @ (posedge clk_SCL_n or negedge dev_rst_n) 
      if (!dev_rst_n)
        st_dev_id    <= 1'b0;
      else if (is_i3c | cf_SlvNack) begin
        if (st_dev_id)
          st_dev_id  <= 1'b0;
      end else if (~st_dev_id & (state==ST_ACK_NACK) & dev_id & ~was_read) 
        st_dev_id    <= isdevid;        // start of Device ID
      else if (st_dev_id & (state==ST_W9TH))
        st_dev_id    <= ourid;          // byte after DevID request is our address
      else if (st_dev_id & (state==ST_ACK_NACK))
        st_dev_id    <= dev_id & was_read; // cancel if not next thing being read

    always @ (posedge clk_SCL or negedge dev_rst_n) 
      if (!dev_rst_n)
        dev_cnt <= 2'd0;       
      else if (~st_dev_id) begin
        if (|dev_cnt)
          dev_cnt <= 2'd0;       
      end else if ((state==ST_ACK_NACK) & was_read & dev_id)
        dev_cnt   <= 2'd3;
      else if (|dev_cnt & (state==ST_R9TH))
        dev_cnt   <= dev_cnt[1] ? (dev_cnt-2'd1) : 2'd3;  // note wraparound
    assign spc_data_valid = |dev_cnt;
    assign spc_data_end   = 0; //dev_cnt==2'd1;
    assign spc_data       = &dev_cnt   ? MAP_I2CID[23:16] : 
                            dev_cnt[1] ? MAP_I2CID[15:8] :
                            {MAP_I2CID[7:3],MAP_I2CID[2:1]|i2c_dev_rev};
  end else begin
    assign dev_spc        = 1'b0;
    assign spc_data_valid = 1'b0;
    assign spc_data_end   = 1'b0;
    assign spc_data       = 8'd0;
  end

  endgenerate

  // possible address match - only looked at at end of address header
  assign is_dasa        = in_ccc[ 7:5] ==  3'd5; // generic DASA
  assign addr_match     = matched_7e |
                          (matched_p2p & is_dasa) | // point to point
                          (is_i3c ? (map_da_match|(work_datab[7:1]==dyn_addr[7:1])|vgpio_spc) : 
                            (map_sa_match|map_spc|dev_spc|(opt_static_addr[0]&(work_datab[7:1]==opt_static_addr[7:1]))));
  wire [7:1] err_tst    = work_datab[7:1] ^ 7'h7E;
  assign s0_error       = (state==ST_A7_A0_RnW) & ~|bit_cnt &  
                          ~matched_7e & ~|(err_tst & (err_tst-1)) &
                          is_i3c; // only S0 error if have a DA
  assign s0_error2      = (state==ST_A7_A0_RnW) & ~|bit_cnt &
                          matched_7e & SDA_r & ~in_daa_mode; // 7E/R out of DAA
  wire   is_da_matched  = (state==ST_ACK_NACK) & is_i3c & ~cf_SlvNack & 
                          (map_da_match|(work_datab[7:1]==dyn_addr[7:1])) & 
                          (~ccc_handled | ccc_valid_dir);
  wire   da_suppress    = vgpio_ccc | (in_ccc[0] & in_ccc[ 2] & ~ccc_uh_r & ~ccc_handled); // these collapse if not used
  assign int_da_matched = is_da_matched & ~da_suppress;
  assign int_sa_matched = (state==ST_ACK_NACK) & ~is_i3c & ~cf_SlvNack &
                          ((map_sa_match&~dev_spc)|(opt_static_addr[0]&(work_datab[7:1]==opt_static_addr[7:1])));
  assign matched_7e     = (work_datab[7:1]==7'h7E); // not ~cf_SlvNack as 7E must be ACKed
  assign matched_p2p    = (work_datab[7:1]==7'h01);
  assign int_7e_matched = (state==ST_ACK_NACK) & matched_7e; // pulse
    // note on below: for SETGET types, ccc_get matches was_read
  assign ccc_valid_dir  = ccc_get == was_read;
  assign ccc_is_read    = (state==ST_A7_A0_RnW) ? SDA_r : was_read;
  wire   is_our_ccc     = state_in_CCC[ 1] | (is_dasa & int_sa_matched) |
                          is_da_matched;
  // ack_ok is controlled based on context. If our dynamic or static addr,
  // then we suppress if buffer not ready. If 7E, then we normally ACK
  // except for ENTDAA's 7E/R once we have DA. Note that tb_data_valid is
  // a risk because it can go from empty to ~empty mid-cycle, so we reg
  reg    was_tb_val;
  wire   is_tb_val      = (tb_data_valid & was_tb_val) | dev_spc; 
  wire   valid_sa       = int_sa_matched & (~in_ccc[ 2] | is_dasa);
  wire   our_data       = is_da_matched | valid_sa | map_spc | dev_spc | vgpio_spc;
  wire   our_ready      = ccc_handled ? is_our_ccc : 
                                        (was_read ? is_tb_val : fb_data_use);
  wire   daa_7e         = (ccc_byte[7:1]==7'h7E);
  assign ack_ok         = in_daa_mode ? (was_read&daa_7e&~dyn_addr[0]) : // check for S4
                          our_data ? our_ready : 
                          ((matched_7e & ~was_read) |
                           (matched_p2p & is_dasa)); // point to point
  always @ (posedge clk_SCL or negedge RSTn) 
    if (!RSTn)
      was_tb_val <= 1'b1;
    else if (our_data)
      was_tb_val <= tb_data_valid;      // mask for ack OK over SCL rising
    else if (~was_tb_val)
      was_tb_val <= 1'b1;

  // next allows pad to enable i2c spike filter until i3c
  // we have to use the 7e check so disables in ACK/NACK cycle
  reg spike_lock;
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      spike_lock <= 1'b0;
    else if (~spike_lock)
      if ((state[2:0] == ST_A7_A0_RnW[2:0]) & (bit_cnt == 3'h0) & (work_datab[7:1]==7'h7E))
        spike_lock <= 1'b1;               // I3C, never allow spike filter
      else if (is_i3c)
        spike_lock <= 1'b1;               // if I3C DA valid, no spike
  assign i2c_spike_ok = ~spike_lock;

  // the event_won reg allows us to handle IBI/MR/HJ during start
  // it gets set on falling edge of A7 if we want to arb
  // it will get cleared if we "lose" an arb bit (A7 to A0 to RnW)
  assign event_losing = SDA_r != event_hdr_b;
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      event_won   <= 1'b0;
    else if (event_pending[2])          // if we want to inject event - happens on START
      event_won   <= 1'b1;              // resets on STOP
    else if (|event_pending[1:0]) begin
      if (state == ST_EV_A7_A0_RnW) begin
        if (event_losing)               // clear if loses 
          event_won <= 1'b0;
      end else if ((next_state != ST_MACK) & (state != ST_IBI_BYTE) & (next_state != ST_IBI_BYTE))
        event_won <= 1'b0;              // clear if not in that header anymore
    end else
      event_won   <= 1'b0;

  // next is just for IBI extended data (past MDB). If not supported,
  // this will all go away since ibi_extdata would be const 0
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      ibi_rd      <= 1'b0;
    else if ((state==ST_MACK) & pend_IBI & ibi_has_byte & |extdata_pos) 
      ibi_rd      <= 1'b1;              // data read past 1st byte
    else if (ibi_rd)
      if ((state==ST_WAIT_SrP) | (state==ST_A7_A0_RnW))
        ibi_rd    <= 1'b0;              // clear on Sr or P (or next S)
  assign ibi_in_extdata = ibi_rd & ~int_in_STOP;

  // event sent/ack tell the upper layers if we are done. We have to do
  // in MACK (master ack) if NACK or if no IBI byte and ACK. But, if 
  // and IBI byte and ACK, then we have to wait until we are past the
  // MACK state or we could get canceled before we transition.
  // note below: ack using live pin_SDA_in is safe as it is registered
  // in upper layers using SCL (same as SDA_in). 
  // The use of the SDA pin is because ack is registered on SCL rising
  assign int_event_sent = ((state==ST_MACK) & 
                           (pin_SDA_in | ~ibi_has_byte | ~pend_IBI)) | // nack or no-byte
                          ((state==ST_IBI9TH) & pend_IBI); // pend_IBI clears after
  assign int_event_ack  = ((state==ST_MACK) & ~pin_SDA_in) |
                           (state==ST_IBI9TH);
  assign int_in_evproc  = (state==ST_MACK) | (state==ST_EV_A7_A0_RnW) |
                          (state==ST_IBI_BYTE) | (state==ST_IBI9TH);
    // next two only apply if time stamping used. Note that we have
    // to mark rising edge of 1st bit of IBI byte and then 1st bit
    // of SC1
  assign ibi_timec_marks[0] = (state==ST_IBI_BYTE) & (bit_cnt==3'd7) &
                              ~|extdata_pos;
  assign ibi_timec_marks[1] = (state==ST_IBI_BYTE) & (bit_cnt==3'd7) &
                              (extdata_pos==2'd1);
    // next has no param because ibi_extdata would be const 0 if not supported
  assign ibi_exttx          = ibi_extdata[2] & tb_data_valid;

  //
  // 7. Tracking mechanism related to CCCs. 
  //

  // in_ccc is a special tracker for CCC as mode/state and because we can have CCC 
  // commands before we have an I3C Dynamic Address. So, we have to know to use I3C 
  // rules and not i2c.
  // The bits are used to track mode and state:
  // [CF_POSS]  =1 means we saw 0x7E/W (write) and so possible CCC
  // [CF_BCAST] =1 means we saw a CCC broadcast command. So [DIRECT]=0. Clears on any START
  // [CF_DIRECT]=1 means we saw a CCC direct command.    So [BCAST]=0.  Clears on 7E/W and START
  // [CF_DAA_M] =1 means we saw ENTDAA command and so in mode.          Clears on STOP only
  // [CF_HDR]   =1 means we saw ENTHDR command and so in mode.          Clears on HDR Exit
  // [CF_GRP_b] =1 to 7 for specific commands, else 0 if none. 
  // Note that S5 is unclear - we ignore unknown CCCs, so likely enough
  // Note: we are using fb_datab_r in W9TH to allow for parity errors
  // CCC cannot be valid past STOP, so we reset since no clock for stop
  wire ccc_reset_n = RSTn & (~STOP_r | scan_no_rst); 
  always @ (posedge clk_SCL_n or negedge ccc_reset_n) 
    if (!ccc_reset_n)
      in_ccc                <= 5'd0;
    else if ((state==ST_ACK_NACK) & ~was_read) begin
      // saw (re)START,addr/W -- so see if for us if direct, else clear if 7E
      if (matched_7e) begin
        // 7E on direct clears CCC state
        if (in_ccc[ 2] | in_ccc[ 3]) // either S4 error or natural end
          in_ccc[ 4:0]<= 5'd1;    // clear direct and DAA with 7E/W
        else
          in_ccc[ 0]  <= 1'b1;    // always true even if already set
        in_ccc[ 7:5]   <= 3'd0;
      end else if (is_dasa) begin       // SETDASA - see if matches us
        // This is special handled because DASA not using our dynamic address
        // note that Mapped SAs not used here - only base SA
        if (opt_static_addr[0] & (work_datab[7:1]==opt_static_addr[7:1]))
          in_ccc[ 7:5] <=  3'd6; // matches our SA
        else if (work_datab[7:1]==7'h01)
          in_ccc[ 7:5] <=  3'd7;// matches point to point
        else
          in_ccc[ 7:5] <=  3'd5;    // matches none
      end else if (in_ccc[ 7:5] ==  3'd2) begin
        if (is_i3c & (work_datab[7:1]==dyn_addr[7:1]))
          in_ccc[ 7:5] <=  3'd3;
      end
    end else if ((in_ccc[ 2:0]==3'b001) & (state==ST_W9TH) & (parity==SDA_r)) begin
      // was possible and now we have a command value
      // 1st byte is command. If bit [7]=1, then direct, else broadcast
      in_ccc[ 2: 1] <= work_datab[7] ? 2'b10 : 2'b01;
      if (work_datab ==  8'h07)
        in_ccc[ 3]   <= 1'b1;     // entered DAA mode
      else if (work_datab[7:3] ==  (8'h20>>3)) begin
        in_ccc[ 4]   <= 1'b1;     // entered HDR mode
        // for actual handling, we only store DDR
        in_ccc[ 7:5]   <= |work_datab[2:0] ? 3'd0 :  3'd4;
      end else if (work_datab ==  8'h87)
        in_ccc[ 7:5]   <=  3'd5;
      else if (work_datab ==  8'h88)
        in_ccc[ 7:5]   <=  3'd1;
      else if (work_datab ==  8'h06)
        in_ccc[ 7:5]   <=  3'd3;
      else if (work_datab ==  8'h86)
        in_ccc[ 7:5]   <=  3'd2; // pending since direct
      else
        in_ccc[ 7:5]   <= 0;
    end else if (|in_ccc & any_START_d) begin
      // STOP-START clears all. repeated START clears possible and also Broadcast if
      // not DAA. Note that HDR would prevent start detect, so this would be real
      if (real_START_d |                // START (after STOP) clears all
          (in_ccc[ 4:0] ==  5'd1) |   // possible only is cleared by repeated START
          (in_ccc[ 1] & ~in_ccc[ 3])) begin // broadcast but not DAA
        in_ccc[ 4:0]  <= 5'd0;    // back to none
        in_ccc[ 7:5]   <= 3'd0;    // clear command if any
      end
    end 
  assign state_in_CCC = in_ccc & {8{~STOP_r}}; // mask when stopped
  assign in_daa_mode  = in_ccc[ 3];
    // next_ccc allows inspection early of the Cmd. Note that parity
    // not checked, but if bad parity we will abandon anyway. We have
    // to resolve early so ccc_handling known - so don't store for app
  assign next_ccc     = ((in_ccc[ 2:0]==3'b001) & 
                          ((next_state==ST_W9TH) | (state_base==ST_W9TH))) ?
                          (work_datab[7] ? 2'b10 : 2'b01) : 2'b00;
  assign next_hdr[0]  = (in_ccc[ 2:0]==3'b001) & (state==ST_W9TH) &
                        (parity==SDA_r) & (work_datab[7:3]== (8'h20>>3));
  assign next_hdr[1]  =  ~|work_datab[2:0];
  generate if (ENA_HDR[ 0]) begin : ddr_act
    reg  ddr_active_r;          // tracks just DDR
    // next is to track entry to DDR but early enough to allow DDR state machine
    // to get started when it is needed. This is tricky due to parity. So, we
    // do on rising edge of SCL.
    always @ (posedge clk_SCL or negedge ccc_reset_n) 
      if (!ccc_reset_n)
        ddr_active_r <= 1'b0;
      else if ((in_ccc[ 2:0]==3'b001) & (state==ST_W9TH) &
               (work_datab[7:3]== (8'h20>>3)) & ~|work_datab[2:0] &
               (parity==pin_SDA_in))
        ddr_active_r <= 1'b1;           // DDR and parity good
    assign in_ddr = ddr_active_r;
  end else begin
    assign in_ddr = 1'b0;
  end endgenerate

  // S4 error is during DAA and cleared on STOP
  always @ (posedge clk_SCL_n or negedge ccc_reset_n) 
    if (!ccc_reset_n)
      s4_err <= 1'b0;
    else if (state==ST_ACK_NACK)
      if (in_daa_mode & (~was_read | ~daa_7e))
        s4_err <= 1'b1;

  // int_ccc is a bit tricky. We want to interrupt only when we have to,
  // but this presents some issues. 
  // 1. If broadcast CCC and not handled, then we want to feed to app. 
  // 2. If direct CCC which is a "SET" and not handled, then we only want 
  //    to tell app if/when we are selected; we then feed saved CCC
  //    to app. 
  // 3. If direct CCC which is a "GET" and not handled, then we have to
  //    tell app from broadcast, but they do not know if they will be 
  //    selected and so would have to prepare an answer???
  assign next_int_ccc = (in_ccc[ 2:0]==3'b001) & (state==ST_W9TH) & 
                        (parity==SDA_r) & ~ccc_handling & ccc_uh_mask & ~cf_SlvNack;
  assign next_bc_ccc  = next_int_ccc & ~work_datab[7];
  assign next_dc_ccc  = ccc_supp[1] & int_da_matched & in_ccc[ 2];
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn) begin
      ccc_supp    <= 2'b00;
      ccc_uh_r    <= 1'b0;              // masked unhandled: it is 1 for the CCC
    end else if (next_int_ccc) begin    // only if not handled
      if (work_datab[7])
        ccc_supp  <= 2'b10;             // only notify on match
      else if (work_datab[7:3] !=  (8'h20>>3)) // exception is HDR
        ccc_supp  <= 2'b01;             // notify now and also push CCC to RX
      ccc_uh_r    <= 1'b1;
    end else if (ccc_supp[0])
      ccc_supp[0] <= 1'b0;              // pulsed once
    else if (ccc_handling | ~ccc_uh_r)
      ccc_supp    <= 2'b00;             // is handled
    else if (next_dc_ccc)
      ccc_supp[0] <= 1'b1;              // signal now (pulse)
    else if (~|in_ccc[ 2:0] & (ccc_supp[1] | ccc_uh_r)) begin
      ccc_supp[1] <= 1'b0;              // no longer possibly Direct
      ccc_uh_r    <= 1'b0;              // end of CCC, so clear
    end
  assign int_ccc = ccc_supp[0] & ~vgpio_ccc;
  always @ (posedge clk_SCL_n or negedge ccc_reset_n) 
    if (!ccc_reset_n)
      pass_ccc   <= 1'b0;
    else if (next_dc_ccc  |             // we emit even if NACKed so they can do 2nd time (GET) 
             (next_bc_ccc & (work_datab[7:3] !=  (8'h20>>3))))
      pass_ccc   <= 1'b1;               // pass along command value if unhandled and not HDR
    else if (pass_ccc)
      pass_ccc   <= 1'b0;

  // daa_lost tracks if we lost arb. We will try again on next round
  generate if (PIN_MODEL ==  0) begin : combo_for_daa
    always @ (posedge clk_SCL_n or negedge RSTn) 
      if (!RSTn)
        daa_lost   <= 1'b0;
      else if (in_daa_mode)
        if (state == ST_ACK_NACK)
          daa_lost <= 1'b0;
        else if ((state == ST_DAA) & daa_inp_drv)
          daa_lost <= ~SDA_r & daa_inp_bit;   // OD arb, so we lose if we have 1, we read 0
    assign daa_active = (state == ST_DAA) & ~is_i3c; // participate if we need a DA
  end else begin
    // registered pad control
    always @ (posedge clk_SCL_n or negedge RSTn) 
      if (!RSTn)
        daa_lost   <= 1'b0;
      else if (in_daa_mode)
        if (state == ST_A7_A0_RnW)
          daa_lost <= 1'b0;
        else if ((state == ST_DAA) & daa_inp_drv)
          daa_lost <= ~SDA_r & ~r_SDA_oena;   // OD arb, so we lose if we have 1, we read 0
    wire next_daa = (state==ST_ACK_NACK) & in_daa_mode & was_read;
    wire stay_daa = (state==ST_DAA) & ~(daa_lost | daa_done);
    assign daa_active = (next_daa | stay_daa)  & ~is_i3c; // participate if we need a DA
  end endgenerate

  // IBI bytes can come from MDB (1st), then time-control if ibi_timec=1,
  // then extdata if ibi_extdata[0]=1. 
  // we run counter to keep simple
  // will collapse if not used
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      extdata_pos          <= 4'd0;
    else if ((state==ST_IBI_BYTE) & ~|bit_cnt & (ibi_timec|ibi_extdata[0])) begin
      if (~|extdata_pos)
        extdata_pos        <= ibi_timec ? 4'b0001 : 4'b1000;  // 1st extended byte
      else if (ibi_timec & ~extdata_pos[3]) begin
        if (~&extdata_pos[1:0])
          extdata_pos[1:0] <= extdata_pos[1:0] + 2'd1;
        else if (ibi_extdata[0])
          extdata_pos      <= 4'b1000;  // extended feed
        else
          extdata_pos      <= 4'd0;     // nothing more to do
      end else if (extdata_pos[3])
        extdata_pos[2:0]   <= extdata_pos[2:0] + 3'd1;
    end else if ((state != ST_IBI_BYTE) & (state != ST_IBI9TH))
      extdata_pos          <= 4'd0;
  assign ibi_out_sel = {(state==ST_IBI_BYTE)|(state==ST_IBI9TH),extdata_pos};

  // below delays the end mark so not before will gho away if not used
  always @ (posedge clk_SCL_n or negedge RSTn) 
    if (!RSTn)
      ext_end <= 1'b0;
    else if (ext_end != ibi_extdata[1])
      ext_end <= ibi_extdata[1];

  //
  // 8. Pad output control for SDA
  //    -- SDR and i2c only affect SDA, not SCL
  //    -- Both OD and PP (std, as used by i3c)
  //
  // Note: the pad controls are combinatorial, since shallow cones and not
  //       very high speed. This could be changed to registered using PIN_MODEL,
  //       but note restrictions for HDR-DDR on external registered.

  wire [7:1] ibi_da;
  generate  if (ENA_MAPPED[ 0]) begin : ibi_mapped
    genvar i,j;
    wire [(MAP_CNT*8)-1:0] map_regs_ad  = SetMappedDASA[(MAP_CNT*10)-1:(MAP_CNT*2)];
    wire [MAP_CNT:0] da[0:7];
    for (i = 1; i <= MAP_CNT; i = i + 1) begin : make_da
      for (j = 0; j < 8; j = j +1) begin : make_da2
        assign da[7-j][i] = (ibi_mapidx==i) ? map_regs_ad[(((i-1)*8)+7) - j] : 1'b0;
      end
    end
    for (j = 0; j < 8; j = j +1) begin : make_da3
      assign da[j][0] = dyn_addr[j];
    end
    for (j = 1; j < 8; j = j +1) begin : make_ibida
      assign ibi_da[j] = |da[j];
    end
  end else begin
    assign ibi_da = dyn_addr[7:1];
  end endgenerate

  // events related to IBI, MR, HJ. Note start_r holds during 1st header
  assign event_drive = start_r & event_won; // one of pend_HJ | pend_IBI | pend_MR
  assign event_hdr   = pend_HJ ? {7'h02,1'b0} : {ibi_da, pend_IBI};
  assign event_hdr_b = event_hdr[bit_cnt];

  // forced START is for forcing IBI and we protect it here for races
  assign force_START = force_sda & (STOP_r | init_stop) & pin_SCL_in;

  // generate the OE/OUT likely values. These will be conditioned below
  generate if (PIN_MODEL ==  0) begin : use_combo_pins
    // combinatorial means we use all registered inputs, but generate the
    // pad out (oe and out) from these moderately shallow cones below. 
    // For process where this depth is fast enough paths, this is used.
    // Else if paths long or path to pads long, then use registered
    assign SDA_oe = 
                  // Header: if not event drive, below is never used. if event_drive, we 
                  // are driving event (until end or we lose); open-drain 
                  ((state==ST_EV_A7_A0_RnW) & ~event_hdr_b) | // event drive used below
                  // ACK/NACK: if event, we are here if we lost. If not event or lost, 
                  // then note I3C ACK on write ends at the rising edge of SCL
                  ((state==ST_ACK_NACK)  & ack_ok &       
                      (~in_i3c_msg | (~event_drive & (~pin_SCL_in | was_read)))) |
                  // ENTDAA: arbitrated push of ID to get assigned an address. Last part
                  // is assignment of our ID. Note that daa_inp_bit changes just after
                  // falling edge of SCL. We ACK new DA on SCL=High
                  ((state==ST_DAA)       & daa_inp_drv & ~daa_lost & ~daa_inp_bit) |
                  // Event Master ACKs/NACKs: special case if ACKed and IBI_byte, 
                  // we have to hold SDA=0 when SCL=1
                  ((state==ST_MACK)      & ((next_state==ST_IBI_BYTE) & pin_SCL_in)) |
                  // if we have a byte after a won IBI, this emits it PP. 
                  (state==ST_IBI_BYTE)     |
                  // IBI 9th is always 0
                  ((state==ST_IBI9TH)    & (|extdata_pos |  ~pin_SCL_in | ibi_exttx)) | 
                  // Read: is same for i2c and i3c except for OD vs. PP
                  ((state==ST_READ)      & ~is_term & (in_i3c_msg | ~rdata_bit)) |
                  // Read 9th: i2c is from master ACK, i3c is T bit
                  // (T high when SCL=0 if not end, else low, then High-Z when SCL=1)
                  ((state==ST_R9TH)      & (in_i3c_msg & ~pin_SCL_in)) | 
                  // Write 9th bit: i2c is us ACKing; i3c input only. 
                  // FUTURE: do we check for overrun in i2c and NACK?
                  ((state==ST_W9TH)      & ~in_i3c_msg) |
                  // if DDR, we change bit depending on edge
                  (ENA_HDR[ 0]   & |ddr_tb_ctrl & ddr_tb_oe[~pin_SCL_in]);

    assign SDA_out = 
                  ((state==ST_DAA)       & daa_inp_bit) |
                  ((state==ST_IBI_BYTE)  & opt_ibi_byte[bit_cnt]) | // special byte
                  ((state==ST_IBI9TH)    & ((|extdata_pos&~ext_end)|ibi_exttx)) | 
                  ((state==ST_READ)      & (in_i3c_msg & rdata_bit)) |
                  ((state==ST_R9TH)      & ~i3c_end) |
                  (ENA_HDR[ 0]   & |ddr_tb_ctrl & rddr_bits[~pin_SCL_in]);// [1] 1st
                   // W9TH for i2c is always ACK (0). 


    // now drive pad for SDA
    wire is_valid_ddr = ENA_HDR[ 0] & |ddr_tb_ctrl;
    assign pin_SDA_oena = force_START | ((|state[3:2] |is_valid_ddr) & SDA_oe);
    assign pin_SDA_out  = pin_SDA_oena ? SDA_out : 1'b1; // 1 to avoid oe/out glitch
    assign pin_SDA_oena_rise = 1'b0;    // not used
    assign r_SDA_oena = 0;              // not used
  end else begin : use_reg_pins

    // may be: PINM_REG or PINM_EXT_REG depending on whether extended to pads
    // This is used when combo path from PIN_COMB is too slow or the
    // path to pad is too slow. This has to use a lot of combo nets to
    // generate the registered output flops, which is less clean.

    // 1st local analoguess of other nets, usually registered
    wire is_da_or_sa    = is_i3c ? (map_da_match|(work_datab[7:1]==dyn_addr[7:1])|vgpio_spc) : 
                                   (map_sa_match|map_spc|dev_spc|(opt_static_addr[0]&(work_datab[7:1]==opt_static_addr[7:1])));
    wire willbe_read    = SDA_r;
    wire willbe_our_ccc = state_in_CCC[ 1] | (is_dasa & is_da_or_sa) |
                          ((ccc_get == willbe_read) & is_da_or_sa);
    wire our_ack        = ccc_handled ? willbe_our_ccc :
                          (willbe_read ? (tb_data_valid|dev_spc) : fb_data_use);
    wire willbe_daa_7e  = (ccc_byte[7:1]==7'h7E);
    wire pre_ack_ok     = in_daa_mode ? (willbe_read&willbe_daa_7e&~dyn_addr[0]) : // check for S4
                          (is_da_or_sa & ~cf_SlvNack) ? our_ack : 
                          ((matched_7e & ~willbe_read) |
                           (matched_p2p & is_dasa)); // point to point


    // daa_will_lose only matters leading into falling edge, so ignore
    // it after falling before rising
    wire daa_will_lose   = (state==ST_DAA) ? (~SDA_r & ~r_SDA_oena & ~daa_acknack) : 
                                             daa_lost;

    // rdata_bit needs to be pre-computed from pre-cnt vs. cnt
    wire [2:0] rdata_bc  = ((state==ST_ACK_NACK)|(state==ST_R9TH)|(state==ST_IBI9TH)) ? 3'd7 : 
                          (bit_cnt-3'd1);
    wire rdata_prebit    = ccc_handled ? ccc_tb_data :
                           spc_data_valid ? spc_data[rdata_bc] : // i2c special
                           tb_data_valid ? tb_datab[rdata_bc] : 
                           1'b1;           // undriven if i2c and just 1 if i3c

    wire [2:0] event_bc  = int_in_STOP ? 3'd7 : (bit_cnt-3'd1);
    wire event_pre_hdr_b = event_hdr[event_bc];
    wire [2:0] ibi_bc    = ((state==ST_MACK)|(state==ST_IBI9TH)) ? 3'd7 : 
                           (bit_cnt-3'd1);
    // special next for START, which can be unexpected
    wire   next_start    = int_in_STOP | (SDA_r & ~pin_SDA_in);

    wire next_extdata_pos0 = (~ibi_timec&~ibi_extdata[0]) |
                           ((state==ST_IBI_BYTE) & ~|bit_cnt & 
                            ((extdata_pos==4'b0011)|ibi_extdata[1]));

    // we use next_state for each since we have to do before the clock
    // Note that ~pin_SCL_in is in rise_chg below
    assign SDA_oe = (|next_state[3:2] & ~next_start & (
                  //
                  // All standard send to bus: ACK/NACK, ENTDAA, READ, R9TH, i2c-W9TH
                  //
                  // ACK/NACK: if event, we are here if we lost. If not event or lost, 
                  // then note I3C ACK on write ends at the rising edge of SCL
                  ((next_state==ST_ACK_NACK)  & pre_ack_ok) |
                  // ENTDAA: arbitrated push of ID to get assigned an address. Last part
                  // is assignment of our ID. Note that daa_inp_bit changes just after
                  // falling edge of SCL. We ACK new DA as well
                  ((next_state==ST_DAA)       & daa_inp_drv & ~daa_will_lose & ~daa_inp_bit) |
                  // Read: is same for i2c and i3c except for OD vs. PP
                  ((next_state==ST_READ)      & (in_i3c_msg | ~rdata_prebit)) |
                  // Read 9th: i2c is from master ACK, i3c is T bit
                  // (T high when SCL=0 if not end, else low, then High-Z when SCL=1)
                  ((next_state==ST_R9TH)      & in_i3c_msg) | 
                  // Write 9th bit: i2c is us ACKing; i3c input only. 
                  // FUTURE: do we check for overrun in i2c and NACK?
                  ((next_state==ST_W9TH)      & ~in_i3c_msg) |
                  //
                  // Now special for IBI/HJ type uses: only if enabled
                  // -- Drives START header and optional IBI byte
                  //
                  // Header: if not event drive, below is never used. if event_drive, we 
                  // are driving event (until end or we lose); open-drain 
                  ((next_state==ST_EV_A7_A0_RnW) & ~event_pre_hdr_b) | // event drive used below
                  // if we have a byte after a won IBI, this emits it PP. 
                  (next_state==ST_IBI_BYTE)     |
                  // IBI 9th is always driven
                  (next_state==ST_IBI9TH)
                  // ST_MACK not here because it is undriven when SCL=Low, and then drive when High
                  )) |
                  (int_in_STOP & ~SDR_hold & event_pending[2] & |event_pending[1:0] & ~event_pre_hdr_b) | // will drive 1st bit
                  // if DDR, we change bit depending on edge, so rise_chg comes in
                  (ENA_HDR[ 0] & ddr_tb_oe[1]);

    assign SDA_out = 
                  ((next_state==ST_DAA)       & daa_inp_bit)  |
                  ((next_state==ST_IBI_BYTE)  & opt_ibi_byte[ibi_bc]) | // special byte
                  ((next_state==ST_IBI9TH)    & (~next_extdata_pos0 | ibi_exttx)) | 
                  ((next_state==ST_READ)      & (in_i3c_msg & rdata_prebit)) |
                  ((next_state==ST_R9TH)      & ~i3c_end)     |
                  (ENA_HDR[ 0]        & |ddr_tb_ctrl & rddr_bits[1]);// [1] 1st
                   // IBI byte is 1 byte, so always T=0 for 1/2 clock
                   // W9TH for i2c is always ACK (0). 

    // this next one determines if SDA changes on rising of SCL for handoff
    wire rise_chg = in_i3c_msg &        // if we have a DA or if 7E used
                    ~any_START_d & ~STOP_r    &
                    ((r_SDA_oena & (    // if 1, then we are ACKing or T
                       ((state_base==ST_ACK_NACK)  & ~was_read) | // ACK to write handoff
                       (state_base==ST_IBI9TH)                  | // IBI T bit (cont or done)
                       (state_base==ST_R9TH)                    ) // Read T bit (cont or done)
                     ) |
                     (~r_SDA_oena & (   // IBI being ACKed: forced to use pin_SDA_in vs. SDA_r
                       ((state_base==ST_MACK) & (~pin_SDA_in & (pend_IBI & ibi_has_byte))) )
                     ) |
                       (ENA_HDR[ 0]   & ddr_tb_oe[0]) // if chg
                     );

    // now the flops. We register OE only on SCL falling and mask on
    // SCL High if needed (to make High-Z)
    // We always register here so we can compare; if ext_reg, then dup ext
    reg   SDA_oena_r;
    assign r_SDA_oena = SDA_oena_r;
    // note reset below includes read_abort
    always @ (posedge clk_SCL_n or negedge state_rst_n) 
      if (!state_rst_n)
        SDA_oena_r <= 1'b0;
      else if (SDA_oe != SDA_oena_r)  // enable check not needed
        SDA_oena_r <= ~SDA_oena_r;

    if (PIN_MODEL ==  1) begin : pinm_reg_local
      // registered flops handled locally

      // now just the output itself (if OE=1). Note that it only
      // changes on falling and is often parked at 0 (for OD),
      // except for HDR-DDR
      reg SDA_out_r;
      always @ (posedge clk_SCL_n or negedge RSTn) 
        if (!RSTn)
          SDA_out_r <= 1'b0;
        else if (SDA_out_r != SDA_out)
          SDA_out_r <= ~SDA_out_r;
      // next is special for HDR-DDR to form same kind of scheme as rise, but for out. It 
      // also is used for forced SDA Low since it is possible to have SDA left high.
      wire hdr_xor = (force_START & SDA_out_r) | 
                     (ENA_HDR[ 0] & |ddr_tb_oe & pin_SCL_in & (rddr_bits[0] != SDA_out_r));
      // note: OE uses an OR, an XOR, and an AND. 
      

        // allows skipping the extra ~STOP_r check from the command line
        assign pin_SDA_oena = force_START | (SDA_oena_r ^ (rise_chg & pin_SCL_in));
      




      assign pin_SDA_out  = SDA_out_r ^ hdr_xor; // XOR goes away if not HDR
      assign pin_SDA_oena_rise = 1'b0;  // not used
    end else begin : pinm_reg_ext
      // now send combo decision to flops external (close to pads) to
      // get split path time delay. The pin_SDA_oena and pin_SDA_out are
      // only used as input to these flops on falling edge. The 
      // pin_SDA_oena_rise is only used on rising edge. Note that we
      // locally buffer same setting so we know what to do (without a
      // path back). Pad side module: i3c_pinm_reg_ext
      // Note that force_START is only when STOPped and clears
      // on SCL falling. We mix with rise so combinatorial effect
      // and so cleared by pad module when SCL is 0.

      assign pin_SDA_oena = SDA_oe;     // next_SDA_oe
      assign pin_SDA_oena_rise[0] = rise_chg;// set rise rule
      if (ENA_HDR[ 0]) begin : hdr_pinext
        // next is special for HDR-DDR to form same kind of scheme as rise, but for out
        // note this is registered when SCL is high, not combo
        assign pin_SDA_oena_rise[3] = ddr_tb_oe[1];
        assign pin_SDA_oena_rise[2] = rddr_bits[0];
      end 
      assign pin_SDA_oena_rise[1] = force_START;
      assign pin_SDA_out  = SDA_out;    // next_SDA_out
    end

  end endgenerate

endmodule
// This is a STUB only
// NOTE: you must replace this with clock-safe (glitch free) cells
//       for muxing and inversion (or XOR).
//       The PD tools can use constraints to treat the output as
//       clock for clock distribution and STA.


module CLOCK_SOURCE_MUX(
  input    use_scan_clock,
  input    scan_clock,
  input    pin_clock,
  output   clock);

  assign clock = use_scan_clock ? scan_clock : pin_clock;
endmodule

module CLOCK_SOURCE_INV(
  input    scan_no_inv,
  input    good_clock,
  output   clock);

  assign clock = ~scan_no_inv ^ good_clock;
endmodule

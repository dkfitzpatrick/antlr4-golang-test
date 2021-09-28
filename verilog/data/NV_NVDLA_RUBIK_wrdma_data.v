// ================================================================
// NVDLA Open Source Project
// 
// Copyright(c) 2016 - 2017 NVIDIA Corporation.  Licensed under the
// NVDLA Open Hardware License; Check "LICENSE" which comes with 
// this distribution for more information.
// ================================================================

// File Name: NV_NVDLA_RUBIK_wrdma_data.v



// ================================================================
// NVDLA Open Source Project
// 
// Copyright(c) 2016 - 2017 NVIDIA Corporation.  Licensed under the
// NVDLA Open Hardware License; Check "LICENSE" which comes with 
// this distribution for more information.
// ================================================================

// File Name: simulate_x_tick.vh













// deprecated tick defines









// formerly recommended tick defines









// newly recommended tick defines
// (-sv parsing is enabled everywhere now, and explicit widths are no longer needed)











module NV_NVDLA_RUBIK_wrdma_data (
      nvdla_core_clk
    , nvdla_core_rstn
    , idata_prdy
    , idata_pvld





    , idata_pd
    , odata_prdy
    , odata_pvld
    , odata_pd
    , pwrbus_ram_pd
    );

// spyglass disable_block W401 -- clock is not input to module
input         nvdla_core_clk;
input         nvdla_core_rstn;
output        idata_prdy;
input         idata_pvld;





input  [255:0] idata_pd;
input         odata_prdy;
output        odata_pvld;
output [255:0] odata_pd;
input  [31:0] pwrbus_ram_pd;

// Master Clock Gating (SLCG)
//
// We gate the clock(s) when idle or stalled.
// This allows us to turn off numerous miscellaneous flops
// that don't get gated during synthesis for one reason or another.
//
// We gate write side and read side separately. 
// If the fifo is synchronous, we also gate the ram separately, but if
// -master_clk_gated_unified or -status_reg/-status_logic_reg is specified, 
// then we use one clk gate for write, ram, and read.
//
wire nvdla_core_clk_mgated_enable;   // assigned by code at end of this module
wire nvdla_core_clk_mgated;               // used only in synchronous fifos
NV_CLK_gate_power nvdla_core_clk_mgate( .clk(nvdla_core_clk), .reset_(nvdla_core_rstn), .clk_en(nvdla_core_clk_mgated_enable), .clk_gated(nvdla_core_clk_mgated) );

// 
// WRITE SIDE
//
// synopsys translate_off




wire wr_pause_rand;  // random stalling

	

	
// synopsys translate_on
wire wr_reserving;
reg        idata_pvld_in;                                // registered idata_pvld
reg        wr_busy_in;                              // inputs being held this cycle?
assign     idata_prdy = !wr_busy_in;
wire       idata_busy_next;                           // fwd: fifo busy next?

// factor for better timing with distant idata_pvld signal
wire       wr_busy_in_next_wr_req_eq_1 = idata_busy_next;
wire       wr_busy_in_next_wr_req_eq_0 = (idata_pvld_in && idata_busy_next) && !wr_reserving;






wire       wr_busy_in_next = (idata_pvld? wr_busy_in_next_wr_req_eq_1 : wr_busy_in_next_wr_req_eq_0)
                              
                            // synopsys translate_off
                            

                            

                            || wr_pause_rand
                            

                            

                            // synopsys translate_on
                             ;


wire       wr_busy_in_int;
always @( posedge nvdla_core_clk or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        idata_pvld_in <=  1'b0;
        wr_busy_in <=  1'b0;
    end else begin
        wr_busy_in <=  wr_busy_in_next;
        if ( !wr_busy_in_int ) begin
            idata_pvld_in  <=  idata_pvld && !wr_busy_in;
        end
        //synopsys translate_off
            else if ( wr_busy_in_int ) begin
        end else begin
            idata_pvld_in  <=   1'bx; 
        end
        //synopsys translate_on

    end
end

reg        idata_busy_int;		        	// copy for internal use
assign       wr_reserving = idata_pvld_in && !idata_busy_int; // reserving write space?


wire       wr_popping;                          // fwd: write side sees pop?

reg  [2:0] idata_count;			// write-side count

wire [2:0] wr_count_next_wr_popping = wr_reserving ? idata_count : (idata_count - 1'd1); // spyglass disable W164a W484
wire [2:0] wr_count_next_no_wr_popping = wr_reserving ? (idata_count + 1'd1) : idata_count; // spyglass disable W164a W484
wire [2:0] wr_count_next = wr_popping ? wr_count_next_wr_popping : 
                                               wr_count_next_no_wr_popping;

wire wr_count_next_no_wr_popping_is_4 = ( wr_count_next_no_wr_popping == 3'd4 );
wire wr_count_next_is_4 = wr_popping ? 1'b0 :
                                          wr_count_next_no_wr_popping_is_4;
wire [2:0] wr_limit_muxed;  // muxed with simulation/emulation overrides
wire [2:0] wr_limit_reg = wr_limit_muxed;
                          // VCS coverage off
assign     idata_busy_next = wr_count_next_is_4 || // busy next cycle?
                          (wr_limit_reg != 3'd0 &&      // check idata_limit if != 0
                           wr_count_next >= wr_limit_reg)  ;
                          // VCS coverage on
assign     wr_busy_in_int = idata_pvld_in && idata_busy_int;
always @( posedge nvdla_core_clk_mgated or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        idata_busy_int <=  1'b0;
        idata_count <=  3'd0;
    end else begin
	idata_busy_int <=  idata_busy_next;
	if ( wr_reserving ^ wr_popping ) begin
	    idata_count <=  wr_count_next;
        end 
        //synopsys translate_off
            else if ( !(wr_reserving ^ wr_popping) ) begin
        end else begin
            idata_count <=  {3{ 1'bx}};
        end
        //synopsys translate_on

    end
end

wire       wr_pushing = wr_reserving;   // data pushed same cycle as idata_pvld_in

//
// RAM
//

reg  [1:0] idata_adr;			// current write address

// spyglass disable_block W484
// next idata_adr if wr_pushing=1
wire [1:0] wr_adr_next = idata_adr + 1'd1;  // spyglass disable W484
always @( posedge nvdla_core_clk_mgated or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        idata_adr <=  2'd0;
    end else begin
        if ( wr_pushing ) begin
            idata_adr <=  wr_adr_next;
        end
    end
end
// spyglass enable_block W484

wire rd_popping;

reg [1:0] odata_adr;          // read address this cycle
wire ram_we = wr_pushing && (idata_count > 3'd0 || !rd_popping);   // note: write occurs next cycle
wire ram_iwe = !wr_busy_in && idata_pvld;  
wire [255:0] odata_pd_p;                    // read data out of ram

wire [31 : 0] pwrbus_ram_pd;

// Adding parameter for fifogen to disable wr/rd contention assertion in ramgen.
// Fifogen handles this by ignoring the data on the ram data out for that cycle.


NV_NVDLA_RUBIK_wrdma_data_flopram_rwsa_4x256 ram (
      .clk( nvdla_core_clk )
    , .clk_mgated( nvdla_core_clk_mgated )
    , .pwrbus_ram_pd ( pwrbus_ram_pd )
    , .di        ( idata_pd )
    , .iwe        ( ram_iwe )
    , .we        ( ram_we )
    , .wa        ( idata_adr )
    , .ra        ( (idata_count == 0) ? 3'd4 : {1'b0,odata_adr} )
    , .dout        ( odata_pd_p )
    );


wire [1:0] rd_adr_next_popping = odata_adr + 1'd1; // spyglass disable W484
always @( posedge nvdla_core_clk_mgated or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        odata_adr <=  2'd0;
    end else begin
        if ( rd_popping ) begin
	    odata_adr <=  rd_adr_next_popping;
        end 
        //synopsys translate_off
            else if ( !rd_popping ) begin
        end else begin
            odata_adr <=  {2{ 1'bx}};
        end
        //synopsys translate_on

    end
end

//
// SYNCHRONOUS BOUNDARY
//


assign wr_popping = rd_popping;		// let it be seen immediately

wire   rd_pushing = wr_pushing;		// let it be seen immediately

//
// READ SIDE
//

wire       odata_pvld_p; 		// data out of fifo is valid

reg        odata_pvld_int;	// internal copy of odata_pvld
assign     odata_pvld = odata_pvld_int;
assign     rd_popping = odata_pvld_p && !(odata_pvld_int && !odata_prdy);

reg  [2:0] odata_count_p;			// read-side fifo count
// spyglass disable_block W164a W484
wire [2:0] rd_count_p_next_rd_popping = rd_pushing ? odata_count_p : 
                                                                (odata_count_p - 1'd1);
wire [2:0] rd_count_p_next_no_rd_popping =  rd_pushing ? (odata_count_p + 1'd1) : 
                                                                    odata_count_p;
// spyglass enable_block W164a W484
wire [2:0] rd_count_p_next = rd_popping ? rd_count_p_next_rd_popping :
                                                     rd_count_p_next_no_rd_popping; 
assign     odata_pvld_p = odata_count_p != 0 || rd_pushing;
always @( posedge nvdla_core_clk_mgated or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        odata_count_p <=  3'd0;
    end else begin
        if ( rd_pushing || rd_popping  ) begin
	    odata_count_p <=  rd_count_p_next;
        end 
        //synopsys translate_off
            else if ( !(rd_pushing || rd_popping ) ) begin
        end else begin
            odata_count_p <=  {3{ 1'bx}};
        end
        //synopsys translate_on

    end
end
reg [255:0]  odata_pd;         // output data register
wire        rd_req_next = (odata_pvld_p || (odata_pvld_int && !odata_prdy)) ;

always @( posedge nvdla_core_clk_mgated or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        odata_pvld_int <=  1'b0;
    end else begin
        odata_pvld_int <=  rd_req_next;
    end
end
always @( posedge nvdla_core_clk_mgated ) begin
    if ( (rd_popping) ) begin
        odata_pd <=  odata_pd_p;
    end 
    //synopsys translate_off
        else if ( !((rd_popping)) ) begin
    end else begin
        odata_pd <=  {256{ 1'bx}};
    end
    //synopsys translate_on

end

// Master Clock Gating (SLCG) Enables
//

// plusarg for disabling this stuff:

// synopsys translate_off




reg master_clk_gating_disabled;  initial master_clk_gating_disabled = $test$plusargs( "fifogen_disable_master_clk_gating" ) != 0;




// synopsys translate_on

// synopsys translate_off




reg wr_pause_rand_dly;  
always @( posedge nvdla_core_clk or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        wr_pause_rand_dly <=  1'b0;
    end else begin
        wr_pause_rand_dly <=  wr_pause_rand;
    end
end




// synopsys translate_on
assign nvdla_core_clk_mgated_enable = ((wr_reserving || wr_pushing || wr_popping || (idata_pvld_in && !idata_busy_int) || (idata_busy_int != idata_busy_next)) || (rd_pushing || rd_popping || (odata_pvld_int && odata_prdy)) || (wr_pushing))
                               




                               // synopsys translate_off
			       

			       

                               || master_clk_gating_disabled || (wr_pause_rand != wr_pause_rand_dly)
			       

			       

                               // synopsys translate_on
                               ;


// Simulation and Emulation Overrides of wr_limit(s)
//


















 // !EMU

















  

// RTL Simulation Plusarg Override


// VCS coverage off

reg wr_limit_override;
reg [2:0] wr_limit_override_value; 
assign wr_limit_muxed = wr_limit_override ? wr_limit_override_value : 3'd0;















 
initial begin


    wr_limit_override = 0;
    wr_limit_override_value = 0;  // to keep viva happy with dangles
    if ( $test$plusargs( "NV_NVDLA_RUBIK_wrdma_data_wr_limit" ) ) begin
        wr_limit_override = 1;
        $value$plusargs( "NV_NVDLA_RUBIK_wrdma_data_wr_limit=%d", wr_limit_override_value);
    end
end

// VCS coverage on



 






// Random Write-Side Stalling
// synopsys translate_off




// VCS coverage off

// leda W339 OFF -- Non synthesizable operator
// leda W372 OFF -- Undefined PLI task
// leda W373 OFF -- Undefined PLI function
// leda W599 OFF -- This construct is not supported by Synopsys
// leda W430 OFF -- Initial statement is not synthesizable
// leda W182 OFF -- Illegal statement for synthesis
// leda W639 OFF -- For synthesis, operands of a division or modulo operation need to be constants
// leda DCVER_274_NV OFF -- This system task is not supported by DC

integer stall_probability;      // prob of stalling
integer stall_cycles_min;       // min cycles to stall
integer stall_cycles_max;       // max cycles to stall
integer stall_cycles_left;      // stall cycles left




 
initial begin


    stall_probability      = 0; // no stalling by default
    stall_cycles_min       = 1;
    stall_cycles_max       = 10;





    if ( $test$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_probability" ) ) begin
        $value$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_probability=%d", stall_probability);
    end else if ( $test$plusargs( "default_fifo_stall_probability" ) ) begin
        $value$plusargs( "default_fifo_stall_probability=%d", stall_probability);
    end

    if ( $test$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_cycles_min" ) ) begin
        $value$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_cycles_min=%d", stall_cycles_min);
    end else if ( $test$plusargs( "default_fifo_stall_cycles_min" ) ) begin
        $value$plusargs( "default_fifo_stall_cycles_min=%d", stall_cycles_min);
    end

    if ( $test$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_cycles_max" ) ) begin
        $value$plusargs( "NV_NVDLA_RUBIK_wrdma_data_fifo_stall_cycles_max=%d", stall_cycles_max);
    end else if ( $test$plusargs( "default_fifo_stall_cycles_max" ) ) begin
        $value$plusargs( "default_fifo_stall_cycles_max=%d", stall_cycles_max);
    end



    if ( stall_cycles_min < 1 ) begin
        stall_cycles_min = 1;
    end

    if ( stall_cycles_min > stall_cycles_max ) begin
        stall_cycles_max = stall_cycles_min;
    end

end






// randomization globals













always @( negedge nvdla_core_clk or negedge nvdla_core_rstn ) begin
    if ( !nvdla_core_rstn ) begin
        stall_cycles_left <=  0;
    end else begin





            if ( idata_pvld && !(!idata_prdy)
                 && stall_probability != 0 ) begin
                if ( prand_inst0(1, 100) <= stall_probability ) begin
                    stall_cycles_left <=  prand_inst1(stall_cycles_min, stall_cycles_max);
                end else if ( stall_cycles_left !== 0  ) begin
                    stall_cycles_left <=  stall_cycles_left - 1;
                end
            end else if ( stall_cycles_left !== 0  ) begin
                stall_cycles_left <=  stall_cycles_left - 1;
            end


    end
end

assign wr_pause_rand = (stall_cycles_left !== 0) ;

// VCS coverage on




// synopsys translate_on
// VCS coverage on

// leda W339 ON
// leda W372 ON
// leda W373 ON
// leda W599 ON
// leda W430 ON
// leda W182 ON
// leda W639 ON
// leda DCVER_274_NV ON


//
// Histogram of fifo depth (from write side's perspective)
//
// NOTE: it will reference `SIMTOP.perfmon_enabled, so that
//       has to at least be defined, though not initialized.
//	 tbgen testbenches have it already and various
//	 ways to turn it on and off.
//




















// spyglass disable_block W164a W164b W116 W484 W504










// synopsys translate_off

































// synopsys translate_on




























// spyglass enable_block W164a W164b W116 W484 W504


//The NV_BLKBOX_SRC0 module is only present when the FIFOGEN_MODULE_SEARCH
// define is set.  This is to aid fifogen team search for fifogen fifo
// instance and module names in a given design.






// spyglass enable_block W401 -- clock is not input to module

// synopsys dc_script_begin
//   set_boundary_optimization find(design, "NV_NVDLA_RUBIK_wrdma_data") true
// synopsys dc_script_end























function [31:0] prand_inst0;
//VCS coverage off
    input [31:0] min;
    input [31:0] max;
    reg [32:0] diff;
    
    begin

































        prand_inst0 = $RollPLI(min, max, "auto");








    end
//VCS coverage on
endfunction























function [31:0] prand_inst1;
//VCS coverage off
    input [31:0] min;
    input [31:0] max;
    reg [32:0] diff;
    
    begin

































        prand_inst1 = $RollPLI(min, max, "auto");








    end
//VCS coverage on
endfunction


endmodule // NV_NVDLA_RUBIK_wrdma_data

// 
// Flop-Based RAM (with internal wr_reg)
//
module NV_NVDLA_RUBIK_wrdma_data_flopram_rwsa_4x256 (
      clk
    , clk_mgated
    , pwrbus_ram_pd
    , di
    , iwe
    , we
    , wa
    , ra
    , dout
    );

input  clk;  // write clock
input  clk_mgated;  // write clock mgated
input [31 : 0] pwrbus_ram_pd;
input  [255:0] di;
input  iwe;
input  we;
input  [1:0] wa;
input  [2:0] ra;
output [255:0] dout;

NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_0 (.A(pwrbus_ram_pd[0]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_1 (.A(pwrbus_ram_pd[1]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_2 (.A(pwrbus_ram_pd[2]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_3 (.A(pwrbus_ram_pd[3]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_4 (.A(pwrbus_ram_pd[4]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_5 (.A(pwrbus_ram_pd[5]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_6 (.A(pwrbus_ram_pd[6]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_7 (.A(pwrbus_ram_pd[7]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_8 (.A(pwrbus_ram_pd[8]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_9 (.A(pwrbus_ram_pd[9]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_10 (.A(pwrbus_ram_pd[10]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_11 (.A(pwrbus_ram_pd[11]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_12 (.A(pwrbus_ram_pd[12]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_13 (.A(pwrbus_ram_pd[13]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_14 (.A(pwrbus_ram_pd[14]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_15 (.A(pwrbus_ram_pd[15]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_16 (.A(pwrbus_ram_pd[16]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_17 (.A(pwrbus_ram_pd[17]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_18 (.A(pwrbus_ram_pd[18]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_19 (.A(pwrbus_ram_pd[19]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_20 (.A(pwrbus_ram_pd[20]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_21 (.A(pwrbus_ram_pd[21]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_22 (.A(pwrbus_ram_pd[22]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_23 (.A(pwrbus_ram_pd[23]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_24 (.A(pwrbus_ram_pd[24]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_25 (.A(pwrbus_ram_pd[25]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_26 (.A(pwrbus_ram_pd[26]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_27 (.A(pwrbus_ram_pd[27]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_28 (.A(pwrbus_ram_pd[28]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_29 (.A(pwrbus_ram_pd[29]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_30 (.A(pwrbus_ram_pd[30]));
NV_BLKBOX_SINK UJ_BBOX2UNIT_UNUSED_pwrbus_31 (.A(pwrbus_ram_pd[31]));

reg [255:0] di_d;  // -wr_reg

always @( posedge clk ) begin
    if ( iwe ) begin
        di_d <=  di; // -wr_reg
    end
end

































reg [255:0] ram_ff0;
reg [255:0] ram_ff1;
reg [255:0] ram_ff2;
reg [255:0] ram_ff3;

always @( posedge clk_mgated ) begin
    if ( we && wa == 2'd0 ) begin
	ram_ff0 <=  di_d;
    end
    if ( we && wa == 2'd1 ) begin
	ram_ff1 <=  di_d;
    end
    if ( we && wa == 2'd2 ) begin
	ram_ff2 <=  di_d;
    end
    if ( we && wa == 2'd3 ) begin
	ram_ff3 <=  di_d;
    end
end

reg [255:0] dout;

always @(*) begin
    case( ra ) 
    3'd0:       dout = ram_ff0;
    3'd1:       dout = ram_ff1;
    3'd2:       dout = ram_ff2;
    3'd3:       dout = ram_ff3;
    3'd4:       dout = di_d;
    //VCS coverage off
    default:    dout = {256{ 1'bx}};
    //VCS coverage on
    endcase
end


 // EMU

endmodule // NV_NVDLA_RUBIK_wrdma_data_flopram_rwsa_4x256

// emulation model of flopram guts
//























































































 // EMU



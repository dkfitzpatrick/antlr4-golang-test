core4_auto_player(
input logic clk,
input logic rst,
input logic left,
input logic right,
input logic put,
output logic hsync,
output logic vsync,
output logic [3:0] red,
output logic [3:0] green,
output logic [3:0] blue,
output logic win_a,
output logic win_b,
output logic full_panel,
output logic player,
output logic invalid_move,
output logic [6:0] play,
output logic [1:0] panel[5:0][6:0],
output logic [1:0] panel_automated[5:0][6:0]
);

logic [9:0] v_counter;
logic [9:0] h_counter;

// logic [6:0] play;

state_check_update mod1 (clk,rst,left,right,put,panel,panel_automated,play,player,win_a,win_b,full_panel,invalid_move);

reg [11:0] red_blocks[39:0][79:0];
reg [11:0] green_blocks[39:0][79:0];
reg [11:0] red_play[39:0][79:0];
reg [11:0] green_play[39:0][79:0];

sprites_memory mod2 (red_blocks,green_blocks,red_play,green_play);

VGA_sync mod3 (clk,rst,hsync,vsync,v_counter,h_counter);

RGB_output mod4 (clk,rst,v_counter,h_counter,panel,play,player,red_blocks,green_blocks,red_play,green_play,red,green,blue);

endmodule


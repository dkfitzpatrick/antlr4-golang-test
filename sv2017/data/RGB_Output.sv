module RGB_output(
input logic clk,
input logic rst,
input logic [9:0] v_counter,
input logic [9:0] h_counter,
input logic [1:0] panel[5:0][6:0],
input logic [6:0] play,
input logic player,
input reg [11:0] red_player_blocks[40][80],
input reg [11:0] green_player_blocks[40][80],
input reg [11:0] red_play[40][80],
input reg [11:0] green_play[40][80],
output logic [3:0] red,
output logic [3:0] green,
output logic [3:0] blue
);
 
// assign rst = ~ rst; I changed the testbench instead
// ------------- LOGIC FOR PANEL ---------------------
// the input would be a 2d panel array 6x7 unpacked array with 2bits packed, 7 bits play array and 1 bit turn signal
// create some conditions of v_counter and h_counter and act accordingly.

localparam square_dim_h = 80;
localparam square_dim_v = 40;

localparam h_margin = 10;
// localparam h_margin = 45;
localparam v_margin = 25;

// creating the square conditions
// with packed array
// left is the most significant bit
logic [5:0][6:0] square_conditions;
int row;
int col;
always_comb begin
	for (row = 0; row<6; row=row+1) begin
		for (col = 0; col<7; col=col+1) begin
		
			square_conditions[5-row][6-col] = ((h_counter >= h_margin+col*(h_margin+square_dim_h)) && (h_counter <= (h_margin+square_dim_h+col*(h_margin+square_dim_h))) && (v_counter >= v_margin+row*(v_margin+square_dim_v)) && (v_counter <= (v_margin+square_dim_v+row*(v_margin+square_dim_v))))?1:0;
		
		end
	end
end

// ---------------------------LOGIC FOR PLAY AND TURN ------------------
logic [6:0] play_conditions;
int playindex;
localparam play_offset = 6; //in order to use the same code as before...lengths were calculated having 7 rows ..here we use that final row
localparam half_square = 0; //an extra top margin for the play row if needed
// left at 0 right at 6 
always_comb begin
	for (playindex= 0; playindex<7; playindex=playindex+1) begin
		
		play_conditions[6-playindex] = ((h_counter >= h_margin+playindex*(h_margin+square_dim_h)) && (h_counter <= (h_margin+square_dim_h+playindex*(h_margin+square_dim_h))) && (v_counter >= half_square+v_margin+play_offset*(v_margin+square_dim_v)) && (v_counter <= (v_margin+square_dim_v+play_offset*(v_margin+square_dim_v))))?1:0;
		
	end
end


// ------------------------CODE FOR SPRITE FOR PANEL----------------------------------
logic [9:0] counter_initial[5:0][6:0][1:0];
int row2;
int col2;
always_comb begin
	for (row2=0;row2<6;row2=row2+1) begin
		for (col2=0;col2<7;col2=col2+1) begin
					//0 horizontal-col, 1 vertical-row
			counter_initial[5-row2][6-col2][0] = h_margin+col2*(h_margin+square_dim_h);
			counter_initial[5-row2][6-col2][1] = v_margin+row2*(v_margin+square_dim_v);
			
		end
	end
end  


int row_index;
int column_index;
int row3;
int col3;

//helpful block 
always_comb begin
	for (row3=0; row3<6; row3=row3+1)begin
		for (col3=0; col3<7; col3=col3+1)begin
			//based on i and j I access the table counter initial and use its value to subtrack it from h and v counter so I can choose the desired value from the sprite memory
			if (square_conditions[row3][col3]) begin
			
				row_index = v_counter - counter_initial[row3][col3][1];
				column_index = h_counter - counter_initial[row3][col3][0];
				
			end
		end
	end
end




// ------------------------------CODE FOR SPRITE FOR PLAY------------------------
logic [9:0] counter_play_initial[6:0][1:0];
int col4;
always_comb begin
	for (col4=0;col4<7;col4=col4+1) begin
				//0 horizontal-col, 1 vertical-row
		counter_play_initial[6-col4][0] = h_margin+col4*(h_margin+square_dim_h);
		counter_play_initial[6-col4][1] = v_margin+play_offset*(v_margin+square_dim_v);
	end
end  


int column_index_play;
int row_index_play;
int col5;

//helpful block
always_comb begin
	for (col5=0; col5<7; col5=col5+1)begin
		if (play_conditions[col5]) begin
		
			row_index_play = v_counter - counter_play_initial[col5][1];
			column_index_play = h_counter - counter_play_initial[col5][0];
			
		end
	end
end


int i;
int j;
int k;
// black pixel and rst pixel logic
always_ff @(posedge clk or negedge rst) begin

    if (!rst) begin 
		red <= 0;
		green <= 0;
		blue <= 0;
    end
    else begin
        //black pixels CHANGE: IT WAS "h_counter > 640 || v_counter > 480" MADE IT: and thats completly right after an analysis i made 
        if (h_counter >= 640 || v_counter >= 480) begin
			red <= 0;
			green <= 0;
			blue <= 0;
        end
		else begin 
		
			// with this code i remove every problem because 
			// i black all pixels.. the value of the rgb will be the value that is closest 
			// to the end of the block. If no condition is true then it will not get in so it will stay black
			// if it is then it will be closest on the end so the rgb will take this value instead
			red <= 0;
			green <= 0;
			blue <= 0;
			
			// setting the color to zero when no condition is active solved the problem
			for (i=0; i<6; i=i+1)begin
				for (j=0; j<7; j=j+1)begin
					if (square_conditions[i][j]) begin						
						if (panel[i][j]==2'b01) begin
							red <= red_player_blocks[row_index][column_index][3:0];
							green <= red_player_blocks[row_index][column_index][7:4];
							blue <= red_player_blocks[row_index][column_index][11:8];
						end
						else if(panel[i][j]==2'b10) begin
							red <= green_player_blocks[row_index][column_index][3:0];
							green <= green_player_blocks[row_index][column_index][7:4];
							blue <= green_player_blocks[row_index][column_index][11:8];
						end	
					end
				end
			end
			
			// LOGIC FOR PLAY 
			for (k=0; k<7; k=k+1) begin
				if (play[k]) begin
					if(play_conditions[k]) begin
						if(player == 0) begin
							red <= red_play[row_index_play][column_index_play][3:0];
							green <= red_play[row_index_play][column_index_play][7:4];
							blue <= red_play[row_index_play][column_index_play][11:8];
						end
						else begin
							red <= green_play[row_index_play][column_index_play][3:0];
							green <= green_play[row_index_play][column_index_play][7:4];
							blue <= green_play[row_index_play][column_index_play][11:8];
						end
					end 
				end
			end

		end
    end    
end

endmodule


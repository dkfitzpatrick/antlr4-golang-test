module state_check_update(
input logic clk,
input logic rst,
input logic left,
input logic right,
input logic put,
output logic [1:0] panel[5:0][6:0],
output logic [1:0] panel_automated[5:0][6:0],
output logic [6:0] play,
output logic player,
output logic win_a,
output logic win_b,
output logic full_panel,
output logic invalid_move
);

//--------------------------------check state part------------------------------
logic turn;
assign player = turn ; //player A red/01 - 0, player B green/10 - 1

logic [5:0][1:0] horizontal_temp;
logic [1:0] horizontal;
logic [6:0][1:0] vertical_temp;
logic [1:0] vertical;
logic [2:0][1:0] diag1_temp;
logic [1:0] diag1;
logic [2:0][1:0] diag2_temp;
logic [1:0] diag2;
logic [1:0] winner;

logic [5:0][6:0] full_temp;

int i;
int j;
always_comb begin
	//winner logic
	for (i=0;i<6;i=i+1) begin
		
		horizontal_temp[i] =(panel[i][0] & panel[i][1] & panel[i][2] & panel[i][3]) | 
							(panel[i][1] & panel[i][2] & panel[i][3] & panel[i][4]) |
							(panel[i][2] & panel[i][3] & panel[i][4] & panel[i][5]) |
							(panel[i][3] & panel[i][4] & panel[i][5] & panel[i][6]);
	
	end
	
	horizontal  = horizontal_temp[0] | horizontal_temp[1] | horizontal_temp[2] | horizontal_temp[3] | horizontal_temp[4] | horizontal_temp[5];
	// horizontal holds 00 if no horizontal win , 01 if player A/red wins, 10 if player B/green wins, horizontally only 
	for (i=0;i<7;i=i+1) begin
		
		vertical_temp[i] =	(panel[5][i] & panel[4][i] & panel[3][i] & panel[2][i]) | 
							(panel[4][i] & panel[3][i] & panel[2][i] & panel[1][i]) |
							(panel[3][i] & panel[2][i] & panel[1][i] & panel[0][i]);
	
	end
	
	vertical = vertical_temp[0] | vertical_temp[1] | vertical_temp[2] | vertical_temp[3] | vertical_temp[4] | vertical_temp[5] | vertical_temp[6] ;
	
	for (i=0;i<3;i=i+1) begin
		
		diag1_temp[i] =		(panel[i][6] & panel[i+1][5] & panel[i+2][4] & panel[i+3][3]) | 
							(panel[i][5] & panel[i+1][4] & panel[i+2][3] & panel[i+3][2]) |
							(panel[i][4] & panel[i+1][3] & panel[i+2][2] & panel[i+3][1]) | 
							(panel[i][3] & panel[i+1][2] & panel[i+2][1] & panel[i+3][0]);
	
	end
	
	diag1 = diag1_temp[0] | diag1_temp[1] | diag1_temp[2];
	
	for (i=0;i<3;i=i+1) begin
		
		diag2_temp[i] =		(panel[i][3] & panel[i+1][4] & panel[i+2][5] & panel[i+3][6]) | 
							(panel[i][2] & panel[i+1][3] & panel[i+2][4] & panel[i+3][5]) |
							(panel[i][1] & panel[i+1][2] & panel[i+2][3] & panel[i+3][4]) | 
							(panel[i][0] & panel[i+1][1] & panel[i+2][2] & panel[i+3][3]);
	
	end
	
	diag2 = diag2_temp[0] | diag2_temp[1] | diag2_temp[2];
	
	winner = horizontal | vertical | diag1 | diag2;
	
	//a = red, b = green
	if (winner == 2'b01) win_a = 1;
	else if(winner == 2'b10) win_b = 1;
	else begin
		win_a = 0;
		win_b = 0;
	end
	
	// full logic 
	// filling a temp panel with 1 if it has a player value and with 0 if the value is 00 ...then reducing the panel with AND  
	for (i=0; i<6; i=i+1) begin
		for (j=0; j<7; j=j+1) begin
			if (panel[i][j] != 2'b00) full_temp[i][j] = 1'b1;
			else full_temp[i][j] = 0;			
		end
	end
	full_panel = &full_temp;
end 

//---------------------------------automated player logic-----------------------------

// logic [1:0] panel_automated[5:0][6:0];
logic [5:0][1:0] horizontal_temp_auto;
logic [1:0] horizontal_auto;
logic [6:0][1:0] vertical_temp_auto;
logic [1:0] vertical_auto;
logic [2:0][1:0] diag1_temp_auto;
logic [1:0] diag1_auto;
logic [2:0][1:0] diag2_temp_auto;
logic [1:0] diag2_auto;
logic [1:0] winner_auto;

int k;
always_comb begin
	//winner_auto logic

	for (k=0;k<6;k=k+1) begin
		
		horizontal_temp_auto[k] = (panel_automated[k][0] & panel_automated[k][1] & panel_automated[k][2] & panel_automated[k][3]) | 
								  (panel_automated[k][1] & panel_automated[k][2] & panel_automated[k][3] & panel_automated[k][4]) |
								  (panel_automated[k][2] & panel_automated[k][3] & panel_automated[k][4] & panel_automated[k][5]) |
								  (panel_automated[k][3] & panel_automated[k][4] & panel_automated[k][5] & panel_automated[k][6]);
	
	end
	
	horizontal_auto  = horizontal_temp_auto[0] | horizontal_temp_auto[1] | horizontal_temp_auto[2] | horizontal_temp_auto[3] | horizontal_temp_auto[4] | horizontal_temp_auto[5];
	// horizontal_auto holds 00 if no horizontal_auto win , 01 if player A/red wins, 10 if player B/green wins, horizontal_autoly only 
	
	for (k=0;k<7;k=k+1) begin
		
		vertical_temp_auto[k] =	(panel_automated[5][k] & panel_automated[4][k] & panel_automated[3][k] & panel_automated[2][k]) | 
								(panel_automated[4][k] & panel_automated[3][k] & panel_automated[2][k] & panel_automated[1][k]) |
								(panel_automated[3][k] & panel_automated[2][k] & panel_automated[1][k] & panel_automated[0][k]);
	
	end
	
	vertical_auto = vertical_temp_auto[0] | vertical_temp_auto[1] | vertical_temp_auto[2] | vertical_temp_auto[3] | vertical_temp_auto[4] | vertical_temp_auto[5] | vertical_temp_auto[6] ;
	
	
	for (k=0;k<3;k=k+1) begin
		
		diag1_temp_auto[k] =(panel_automated[k][6] & panel_automated[k+1][5] & panel_automated[k+2][4] & panel_automated[k+3][3]) | 
							(panel_automated[k][5] & panel_automated[k+1][4] & panel_automated[k+2][3] & panel_automated[k+3][2]) |
							(panel_automated[k][4] & panel_automated[k+1][3] & panel_automated[k+2][2] & panel_automated[k+3][1]) | 
							(panel_automated[k][3] & panel_automated[k+1][2] & panel_automated[k+2][1] & panel_automated[k+3][0]);
	
	end
	
	diag1_auto = diag1_temp_auto[0] | diag1_temp_auto[1] | diag1_temp_auto[2];
	
	for (k=0;k<3;k=k+1) begin
		
		diag2_temp_auto[k] =(panel_automated[k][3] & panel_automated[k+1][4] & panel_automated[k+2][5] & panel_automated[k+3][6]) | 
							(panel_automated[k][2] & panel_automated[k+1][3] & panel_automated[k+2][4] & panel_automated[k+3][5]) |
							(panel_automated[k][1] & panel_automated[k+1][2] & panel_automated[k+2][3] & panel_automated[k+3][4]) | 
							(panel_automated[k][0] & panel_automated[k+1][1] & panel_automated[k+2][2] & panel_automated[k+3][3]);
	
	end
	
	diag2_auto = diag2_temp_auto[0] | diag2_temp_auto[1] | diag2_temp_auto[2];
	
	winner_auto = horizontal_auto | vertical_auto | diag1_auto | diag2_auto;
end 


int put_row_auto;
logic found_auto; // usage in finding the first open spot on panel
logic invalid_move_auto;
int auto_col_current;
int l; 
always_comb begin
	found_auto = 0;
	// now I need to find the first open row for the put_col column if l dont find open spot its an invalid move
	for (l=0; l<6; l=l+1)begin  //starting from row 5 : the lowest
		if (panel[l][auto_col_current]==2'b00 && !found_auto) begin
			put_row_auto = l;
			found_auto = 1;
		end
	end
	invalid_move_auto = 0;
	if (found_auto == 0) invalid_move_auto = 1;
end

logic reg1,reg2,reg3;
// psudorandom generator, reg3 is the most significang bit even though den exei na leei 
always_ff @(posedge clk, negedge rst) begin
	if (!rst) begin
		reg3 <= 0;
		reg2 <= 0;
		reg1 <= 1;
	end
	else begin 
		reg1 <= reg2;
		reg2 <= reg3;
		reg3 <= reg2^reg1;
	end
end

logic [2:0] random_col;
assign random_col = {reg3,reg2,reg1};


//---------------------------------left/right/put events------------------------------

logic left_reg,left_event,right_reg,right_event,put_reg,put_event;


always_ff @(posedge clk, negedge rst) begin
	if (!rst) begin
		left_reg <= 1'b0;
	end 
	else begin 
		left_reg <= left;
	end
end 

assign left_event = left_reg & (~left);  //falling edge
// assign rising_edge = (~left_reg) & left;

always_ff @(posedge clk, negedge rst) begin
	if (!rst) begin
		right_reg <= 1'b0;
	end 
	else begin 
		right_reg <= right;
	end
end 

assign right_event = right_reg & (~right);


always_ff @(posedge clk, negedge rst) begin
	if (!rst) begin
		put_reg <= 1'b0;
	end 
	else begin 
		put_reg <= put;
	end
end 

assign put_event = put_reg & (~put);


//-------------------------------finding open row and invalid move------------------------------
int c;
int put_col;
int put_row;
logic found; // usage in finding the first open spot on panel

always_comb begin
	// I want to know what bit of play vector is at one ..calculating it with combinational logic
	for (c=0;c<7;c=c+1)begin
		if (play[c] == 1) put_col=c; //there is always an ace and only one since we operate with shifting
	end
	found = 0;
	// now I need to find the first open row for the put_col column if c dont find open spot its an invalid move
	for (c=0; c<6; c=c+1)begin  //starting from row 0 : the lowest
		if (panel[c][put_col]==2'b00 && !found) begin
			put_row = c;
			found = 1;
		end
	end
	invalid_move = 0;
	if (found == 0) invalid_move = 1;
end

enum {INITIAL,PLAY_UPDATE,PANEL_UPDATE,CHECK_STATUS,ENDGAME,INITIALIZE_AUTO,BLOCK_OPPONENT,CHECK_MOVE_AUTO,RANDOM_MOVE} state;

always_ff @ (posedge clk, negedge rst) begin
    if(!rst) state <= INITIAL;
	else begin
		case(state) 
            INITIAL: begin
				// panel <= '{default:0};
				panel <='{'{2'b10,2'b00,2'b00,2'b00,2'b00,2'b00,2'b10},
						'{2'b01,2'b01,2'b00,2'b00,2'b00,2'b00,2'b01},
						'{2'b10,2'b10,2'b01,2'b00,2'b00,2'b00,2'b10},
						'{2'b01,2'b01,2'b10,2'b10,2'b00,2'b00,2'b01},
						'{2'b10,2'b10,2'b01,2'b01,2'b01,2'b10,2'b10},
						'{2'b01,2'b01,2'b10,2'b10,2'b01,2'b01,2'b01}};
				// panel_automated <= '{default:0};
				panel_automated <='{'{2'b10,2'b00,2'b00,2'b00,2'b00,2'b00,2'b10},
									'{2'b01,2'b01,2'b00,2'b00,2'b00,2'b00,2'b01},
									'{2'b10,2'b10,2'b01,2'b00,2'b00,2'b00,2'b10},
									'{2'b01,2'b01,2'b10,2'b10,2'b00,2'b00,2'b01},
									'{2'b10,2'b10,2'b01,2'b01,2'b01,2'b10,2'b10},
									'{2'b01,2'b01,2'b10,2'b10,2'b01,2'b01,2'b01}};
				// play <= 7'b1000000;
				play <= 7'b0000100;
				turn <= 0; //red plays first
				state <= PLAY_UPDATE;
			end
			PLAY_UPDATE: begin
				
				if (turn == 0 ) begin
					if(put_event) begin 
						state <= PANEL_UPDATE;
					end
					else if (left_event) begin //circular shift
							play <= play << 1;
							play[0] <= play[6];
							state <= PLAY_UPDATE;
					end
					else if (right_event) begin
							play <= play >> 1;
							play[6] <= play[0];
							state <= PLAY_UPDATE;					
					end
				end
				else if (put_event & left_event & right_event) begin//automated player plays
					state <= INITIALIZE_AUTO;
				end
			end
			PANEL_UPDATE: begin
				if (turn == 0) begin
					if (invalid_move) begin					
					state <= PLAY_UPDATE;
					end
					else begin
					panel[put_row][put_col] <= 2'b01;
					panel_automated[put_row][put_col] <= 2'b01; //keep temp_panel updated
					// turn <= ~turn;
					state <= CHECK_STATUS;
					end
				end
				else begin
					if (invalid_move_auto) begin //embeding invalid check on panel update this is only for random move
						if (auto_col_current == 6) begin 
							auto_col_current <= 0;
							state = PANEL_UPDATE;
						end
						else begin
							auto_col_current <= auto_col_current + 1;
							state <= PANEL_UPDATE;
						end
					end
					else begin  //there will be for sure one valid move...if it isnt then i wouldnt be on this state..so the above wont loop forever unless you initialize panel to be full and let auto player play first which will never happen in normal game conditions (it will also loop if player 0 plays first)
						panel[put_row_auto][auto_col_current] <= 2'b10;
						panel_automated[put_row_auto][auto_col_current] <= 2'b10;
						state <= CHECK_STATUS;
					end
				end
			end
			CHECK_STATUS: begin
				if (win_a || win_b || full_panel) state <= ENDGAME;
				else begin 
					state <= PLAY_UPDATE;
					turn <= ~turn;   //changing player after I update panel ...better with testbench  
				end
			end
			ENDGAME: begin
				state <= ENDGAME; //if I go to INITIAL state I lose the last frame 
			end
			INITIALIZE_AUTO: begin
				auto_col_current <= 0;
				state <= BLOCK_OPPONENT;
			end
			BLOCK_OPPONENT: begin
				if (~invalid_move_auto) begin // ama einai miden tote dexome oti uparxei valid seira gia to sigekrimeno column auto_col
					panel_automated[put_row_auto][auto_col_current] <= 2'b01;
					state <= CHECK_MOVE_AUTO;
				end 
				else begin  //if im on last col and I have invalid move then go on RANDOM_MOVE
					if (auto_col_current == 6) begin 
						state <= RANDOM_MOVE;
					end
					else begin  //if gia tin dedomeni stili exw invalid move, allakse sthli kai ksanadokimase
						auto_col_current <= auto_col_current + 1;
						state <= BLOCK_OPPONENT;  //if panel is full I will not be on this state...
						//If panel is not full I will have at least one move that is not invalid and get past the if condition and go to CHECK_MOVE_AUTO
					end
				end
			end
			CHECK_MOVE_AUTO: begin
				if (winner_auto == 2'b01) begin
					state <= PANEL_UPDATE;  // i dont need to remove the value on panel_automated since ill overwrite it on panel update
				end
				else begin
					if (auto_col_current == 6) begin // If I have checked every col and did not found any win logic to block 
						panel_automated[put_row_auto][auto_col_current] <= 2'b00;
						state <= RANDOM_MOVE;  //Move from logic to block the opponent to logic for a random move
					end
					else begin
						auto_col_current <= auto_col_current + 1; 
						panel_automated[put_row_auto][auto_col_current] <= 2'b00;  //remove the last since it wasnt a winning combination..if it was you wouldn't be here
						state <= BLOCK_OPPONENT;
					end
				
				end
			end
			RANDOM_MOVE: begin
				if (random_col == 3'b111) begin //after I place the value on auto_col_current current invalid_move_auto will tell me if its a valid col or not
					auto_col_current <= 0;
				end 
				else begin
					auto_col_current <= random_col;  // the logic is that I choose an initial random col and then check seiriaka tis upoloipes
				end
				state <= PANEL_UPDATE;  //do the changes and go to PANEL_UPDATE anyway 
			end
		endcase
	end
end

endmodule


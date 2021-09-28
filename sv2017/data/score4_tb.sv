module score4_tb;

// set VGA_LOG = 0 to disable logging of the VGA output
parameter logic VGA_LOG = 1; //change on task move as well

// TESTBENCH VARIABLES
integer fileout;
integer frame_cnt;
logic write_frame;
string s;

enum {WIN_1_D, WIN_1_R, WIN_1_C, WIN_2_D, WIN_2_R, WIN_2_C,
	  H_ERROR, V_ERROR, FULL, MY_GAME
	  } scenario;


// DUT SIGNALS
logic clk;
logic rst;

logic left;
logic right;
logic put;

logic player;
logic invalid_move;
logic win_a;
logic win_b;
logic full_panel;

logic vga_hsync;
logic vga_vsync;
logic [3:0] vga_red;
logic [3:0] vga_green;
logic [3:0] vga_blue;

logic [1:0] panel[5:0][6:0];
logic [1:0] panel_automated[5:0][6:0];
logic [6:0] play;
// DUT INSTANTIATION
score4_auto_player DUT(
	.clk   		  (clk),
	.rst   		  (rst),
	.left  		  (left),
	.right 		  (right),
	.put   		  (put),
	.player       (player),
	.invalid_move (invalid_move),
	.win_a 		  (win_a),
	.win_b 		  (win_b),
	.full_panel   (full_panel),
	.hsync 		  (vga_hsync),
	.vsync 		  (vga_vsync),
	.red   		  (vga_red),
	.green 		  (vga_green),
	.blue  		  (vga_blue),
	.panel		  (panel),
	.panel_automated (panel_automated),
	.play			(play)
);


// 50MHz CLOCK GENERATION
always begin
	clk = 0;
	#10ns;
	clk = 1;
	#10ns;
end


// STIMULUS
initial begin
	$timeformat(-9, 0, " ns", 6);

	RESET();
	//$display("\n\n~~~~~~~~ GAME STARTS ~~~~~~~~");
		
	//select one of the following scenarios:
	//	WIN_1_D: player 1 wins (diagonal)
	//	WIN_1_R: player 1 wins (row)
	//	WIN_1_C: player 1 wins (column)
	//	WIN_2_D: player 2 wins (diag)
	//  WIN_2_R: player 2 wins (row)
	//	WIN_2_C: player 2 wins (column)
	//  H_ERROR: check for errors on rows
	//  V_ERROR: check for errors on columns
	//  FULL:    Fill the panel -- WARNING! this creates 70 frames ~> almost 2GB
	//  MY_GAME: create your own game 
	scenario = MY_GAME;
	
	case(scenario)
	MY_GAME: begin 
		
		// YOUR GAME HERE
		// Choose your moves: go_right(), go_left(), lets_put()
			$display("initial panel is");
			$display("%p",panel[5]);
			$display("%p",panel[4]);
			$display("%p",panel[3]);
			$display("%p",panel[2]);
			$display("%p",panel[1]);
			$display("%p",panel[0]);
			$display("\n");
			
			lets_put();
			//im on play update
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			automated_player();
			go_left();
			lets_put();
			
			$display("last panel print");
			$display("%p",panel[5]);
			$display("%p",panel[4]);
			$display("%p",panel[3]);
			$display("%p",panel[2]);
			$display("%p",panel[1]);
			$display("%p",panel[0]);
			$display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b\n", player, invalid_move, win_a, win_b, full_panel);

	
	
		end // end MY_GAME
		
	WIN_1_D:
		p1_wins_d();
		
	WIN_1_R:
		p1_wins_r();
		
	WIN_1_C:
		p1_wins_c();
	
	WIN_2_D: 
		p2_wins_d();
		
	WIN_2_R: 
		p2_wins_r();
	
	WIN_2_C: 
		p2_wins_c();
		
	H_ERROR: 
		check_h_error();
		
	V_ERROR: 
		check_v_error();
		
	FULL: 
		check_full();
	
	default 
		$display("You didn't choose a valid scenario!");
	endcase
	
	
	$display("~~~~~~~~~ GAME OVER ~~~~~~~~~");
		
	$finish;
end

// RESET TASK
task RESET();
	scenario <= MY_GAME;
	write_frame <= 0;
	frame_cnt <= 0;
	
	left <= 0;
	right <= 0;
	put <= 0;
	
	rst <= 0;
	repeat(2) @(posedge clk);
	rst <= 1;
	$display("\n\n~~~~~~~~ GAME STARTS ~~~~~~~~");
	repeat(10) @(posedge clk);
endtask


// Choose action: go left
task go_left();
	$display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b", player, invalid_move, win_a, win_b, full_panel);	
	$display("Moving Left...");
	make_move(1'b0, 1'b0, 1'b1);
endtask


// Choose action: go right
task go_right();
	$display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b", player, invalid_move, win_a, win_b, full_panel);
	$display("Moving Right...");
	make_move(1'b0, 1'b1, 1'b0);
endtask


// Choose action: put
task lets_put();
	$display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b", player, invalid_move, win_a, win_b, full_panel);
	$display("Placing Token!");
	
	make_move(1'b1, 1'b0, 1'b0);
	
	// if (win_a == 1 || win_b == 1 || full_panel == 1) $display("the final panel is");
	// else $display("the panel is");
	// $display("%p",panel[5]);
	// $display("%p",panel[4]);
	// $display("%p",panel[3]);
	// $display("%p",panel[2]);
	// $display("%p",panel[1]);
	// $display("%p",panel[0]);
	// $display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b\n", player, invalid_move, win_a, win_b, full_panel);
endtask

// Choose action: do_nothing
task do_nothing();
	
	$display("Nothing for this frame.");
	make_move(1'b0, 1'b0, 1'b0);
endtask


// Select the next action
task make_move(input logic p, r, l);	

	put <= p;
	right <= r;
	left <= l;
	
	repeat(3) @(posedge clk);
	
	put <= 0;
	right <= 0;
	left <= 0;
	
	
	// comment or uncomment to wait for at least one frame between moves
	write_frame <= 1;  //enable for frame disable for not
	@(negedge write_frame);
	@(posedge clk);
	
	
	
	//the below clocks are good if write frame is not used
	
	// @(posedge clk); // after this edge move_event is up 
	// @(posedge clk); //added //after this clock state will be panel_update / if on play update then left right event is completed
	// @(posedge clk); //added  //after this clock panel updates and state=check_status OR im on play update if invalid move
	// @(posedge clk); //added   if no win then after this clock i go on PLAY_UPDATE
	
	// //either way after these clocks im on play update 
	// // If invalid move Ill wait action from player 0 
	// // else on the next clock I go on INITIALIZE_AUTO
	// //If i press left/right then ill be on play update waiting for action from player 0
	
endtask

task automated_player();
	$display("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b", player, invalid_move, win_a, win_b, full_panel);
	$display("Automated player");
	
	make_move(1'b1, 1'b1, 1'b1);

	// below clocks good if frame capture is not used
	// repeat(25) @(posedge clk);
	// counting for the existing clocks in make move we can set repeat(21) @(posedge clk);
	//after these clocks ill be for sure on PLAY_UPDATE	
	
endtask

// Case study where player 1 wins (diagonal)
task p1_wins_d;
	lets_put();
	go_right();
	
	repeat(2) begin
		lets_put();
		lets_put();
		go_right();
	end
	
	lets_put();
	go_left();
	repeat(2) lets_put();
	go_right();
	repeat(3) lets_put();
endtask


// Case study where player 1 wins (row)
task p1_wins_r;
	repeat(3) begin
		repeat(2) lets_put();
		go_right();
	end
	
	lets_put();
	
endtask


// Case study where player 1 wins (column)
task p1_wins_c;
	repeat(3) begin
		lets_put();
		go_right();
		lets_put();
		go_left();
	end
	lets_put();
	
endtask


// Case study where player 2 wins (diagonal)
task p2_wins_d;
	lets_put();
		
	repeat(2) begin
		go_right();
		lets_put();
	end
		
	lets_put();
		
	go_right();
	repeat(2) lets_put();
		
	go_right();
	lets_put();
		
	go_left();
	lets_put();
		
	go_right();
	repeat(2) lets_put();
		
	go_left();
	lets_put();
	
	go_right();
	lets_put();
endtask


// Case study where player 1 wins (row)
task p2_wins_r;
	repeat(3) begin
		repeat(2) lets_put();
		go_right();
	end
	repeat(2) begin
		go_right();
		lets_put();
		go_left();
		lets_put();
	end

endtask


// Case study where player 2 wins (column)
task p2_wins_c;
	repeat(2) begin
		lets_put();
		go_right();
		lets_put();
		go_left();
	end
	repeat(2) begin
		lets_put();
		go_right();
	end

	lets_put();
	go_left();
	lets_put();
endtask


// Case study for locating errors in a row
task check_h_error;
	lets_put();
	go_left();
	lets_put();
	repeat(8) begin
		go_right();
		lets_put();
	end
		
	go_left();
	lets_put();
		
	go_right();
	lets_put();
endtask


// Case study for locating erros in a column
task check_v_error;
	repeat(7) lets_put();		
	go_right();
	lets_put();
endtask


// Case study for checking a full game panel
task check_full;
	lets_put();
	repeat(6) begin
		go_right();
		lets_put();
	end
	lets_put();
	repeat(6) begin
		go_left();
		lets_put();
	end
	lets_put();
	repeat(6) begin
		go_right();
		lets_put();
	end
	
	repeat(2) begin
		go_left();
		lets_put();
		go_right();
		lets_put();
		lets_put();
		go_left();
		lets_put();
		lets_put();
		go_right();
		lets_put();
	
		repeat(2) go_left();
	end
	
	go_left();
	repeat(3) lets_put();
	go_left();
	repeat(2) lets_put();
	repeat(2) go_right();
	repeat(3) lets_put();
	repeat(2) go_left();
	lets_put();
	
endtask


// MONITORS THE OUTPUTS OF THE DESIGN
initial begin
// $monitor("player : %b , invalid_move : %b , win_a : %b , win_b : %b , full_panel : %b\n", player, invalid_move, win_a, win_b, full_panel);
$monitor("panel_automated is: \n%p\n%p\n%p\n%p\n%p\n%p\n",panel_automated[5],panel_automated[4],panel_automated[3],panel_automated[2],panel_automated[1],panel_automated[0]);
end


// Write frame to VGA log
always @(negedge vga_vsync) begin
	if ( write_frame ) begin
		if (VGA_LOG==1) begin
			s.itoa(frame_cnt);	
			fileout = $fopen({"vga_frame_", s, ".txt"});
		
			repeat (838400) begin
				@(posedge clk);
				$fdisplay(fileout, "%t: %b %b %b %b %b", $time, vga_hsync, vga_vsync, vga_red, vga_green, vga_blue);
			end
			@(negedge clk); 

			frame_cnt ++;
			$fclose(fileout);
		end
		else begin
			repeat (838400) @(posedge clk);
			@(negedge clk);
		end
	$display("frame printed :");
	$display("%p",panel[5]);
	$display("%p",panel[4]);
	$display("%p",panel[3]);
	$display("%p",panel[2]);
	$display("%p",panel[1]);
	$display("%p",panel[0]);
	$display(" {%b, %b, %b, %b, %b, %b, %b}\n",play[6],play[5],play[4],play[3],play[2],play[1],play[0]);
	write_frame <= 0;
	end
	

end



endmodule


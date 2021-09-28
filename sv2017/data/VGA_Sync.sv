module VGA_sync(
input logic clk,
input logic rst,
output logic hsync,
output logic vsync,
output logic [9:0] v_counter,
output logic [9:0] h_counter
);


//pxl_clk acts as an enable signal every 2 clocks of main 50Mhz clk
logic pxl_clk;


// enable/pixel clock logic
always_ff @(posedge clk or negedge rst) begin

    if (!rst) 
        pxl_clk <= 1;
    else 
        pxl_clk <= ~pxl_clk;

end

// horizontal logic
always_ff @(posedge clk or negedge rst) begin

    if (!rst) 
        h_counter <= 0;
    else begin
        if (pxl_clk) begin

            if (h_counter == 799) 
                h_counter <= 0;
            else
                h_counter <= h_counter + 1;
        end    
    end    
end

// vertical logic
always_ff @(posedge clk or negedge rst) begin

    if (!rst) 
        v_counter <= 0;
    else begin
        if (pxl_clk) begin

            if (h_counter == 799) 
                if (v_counter == 523)
                    v_counter <= 0; 
                else
                    v_counter <= v_counter + 1;
        end    
    end    
end

//synchronization
assign hsync = (h_counter > 655 && h_counter < 752)?1'b0:1'b1; //the timings are completly right
assign vsync = (v_counter > 490 && v_counter < 493)?1'b0:1'b1;

endmodule

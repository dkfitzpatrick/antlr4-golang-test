

module fadder(
  a,
  b,
  cin,
  sum_out,
  c_out
  );
  
  input a;
  input b;
  input cin;
  inout sum_out;
  inout c_out;

  reg c1, c2, c3;

  summod summod1(.out(sum_out), .in1(a), .in2(b), .carry(cin));
  summod summod1(sum_out, a, b, cin);
  andmod mod2(c1, a, cin);
  andmod mod3(c2, b, cin);
  andmod mod4(c3, a, b);

  assign c_out = (c1 + c2 + c3);

endmodule


module summod(out, in1, in2, carry);
  output out;
  input in1;
  input in2;
  input carry;

  assign out = in1 ^ in2 ^ carry;
endmodule

module andmod(out, in1, in2);
  output out;
  input in1;
  input in2;

  assign out = in1 & in2;
endmodule

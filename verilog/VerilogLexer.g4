lexer grammar VerilogLexer;

@header {
   import (
      "strconv"
      "strings"
   )
}

@members {
   var CurrentLineNumber int
   var CurrentFileName string
}

channels {
    DEFAULT_CHANNEL,
    HIDDEN_CHANNEL,
    LINE_CHANNEL,
    IGNORED
}

K_module : 'module' ;
K_macromodule : 'macromodule' ;
K_endmodule : 'endmodule' ;
K_B0 : 'B0' ;
K_B1 : 'B1' ;
K_PATHPULSE : 'PATHPULSE' ;
K_always : 'always' ;
K_and : 'and' ;
K_assign : 'assign' ;
K_automatic : 'automatic' ;
K_b0 : 'b0' ;
K_b1 : 'b1' ;
K_begin : 'begin' ;
K_buf : 'buf' ;
K_bufif0 : 'bufif0' ;
K_bufif1 : 'bufif1' ;
K_case : 'case' ;
K_casex : 'casex' ;
K_casez : 'casez' ;
K_config : 'config' ;
K_cell : 'cell' ;
K_cmos : 'cmos' ;
K_deassign : 'deassign' ;
K_default : 'default' ;
K_default_nettype : 'default_nettype' ;
K_defparam : 'defparam' ;
K_design : 'design' ;
K_disable : 'disable' ;
K_edge : 'edge' ;
K_else : 'else' ;
K_end : 'end' ;
K_endcase : 'endcase' ;
K_endconfig : 'endconfig' ;
K_endfunction : 'endfunction' ;
K_endgenerate : 'endgenerate' ;
K_endprimitive : 'endprimitive' ;
K_endspecify : 'endspecify' ;
K_endtable : 'endtable' ;
K_endtask : 'endtask' ;
K_event : 'event' ;
K_for : 'for' ;
K_force : 'force' ;
K_forever : 'forever' ;
K_fork : 'fork' ;
K_highz0 : 'highz0' ;
K_highz1 : 'highz1' ;
K_function : 'function' ;
K_generate : 'generate' ;
K_genvar : 'genvar' ;
K_if : 'if' ;
K_ifnone : 'ifnone' ;
K_include : 'include' ;
K_initial : 'initial' ;
K_inout : 'inout' ;
K_input : 'input' ;
K_instance : 'instance' ;
K_integer : 'integer' ;
K_join : 'join' ;
K_large : 'large' ;
K_liblist : 'liblist' ;
K_library : 'library' ;
K_line : 'line' ;
K_localparam : 'localparam' ;
K_medium : 'medium' ;
K_nand : 'nand' ;
K_negedge : 'negedge' ;
K_nmos : 'nmos' ;
K_nor : 'nor' ;
K_noshowcancelled : 'noshowcancelled' ;
K_not : 'not' ;
K_notif0 : 'notif0' ;
K_notif1 : 'notif1' ;
K_or : 'or' ;
K_output : 'output' ;
K_parameter : 'parameter' ;
K_pmos : 'pmos' ;
K_posedge : 'posedge' ;
K_primitive : 'primitive' ;
K_pull0 : 'pull0' ;
K_pull1 : 'pull1' ;
K_pulldown : 'pulldown' ;
K_pullup : 'pullup' ;
K_pulsestyle_ondetect : 'pulsestyle_ondetect' ;
K_pulsestyle_onevent : 'pulsestyle_onevent' ;
K_rcmos : 'rcmos' ;
K_real : 'real' ;
K_realtime : 'realtime' ;
K_reg : 'reg' ;
K_release : 'release' ;
K_repeat : 'repeat' ;
K_rnmos : 'rnmos' ;
K_rpmos : 'rpmos' ;
K_rtran : 'rtran' ;
K_rtranif0 : 'rtranif0' ;
K_rtranif1 : 'rtranif1' ;
K_scalared : 'scalared' ;
K_showcancelled : 'showcancelled' ;
K_signed : 'signed' ;
K_small : 'small' ;
K_specify : 'specify' ;
K_specparam : 'specparam' ;
K_strong0 : 'strong0' ;
K_strong1 : 'strong1' ;
K_supply0 : 'supply0' ;
K_supply1 : 'supply1' ;
K_table : 'table' ;
K_task : 'task' ;
K_time : 'time' ;
K_timescale : 'timescale' ;
K_tran : 'tran' ;
K_tranif0 : 'tranif0' ;
K_tranif1 : 'tranif1' ;
K_tri : 'tri' ;
K_tri0 : 'tri0' ;
K_tri1 : 'tri1' ;
K_triand : 'triand' ;
K_trior : 'trior' ;
K_trireg : 'trireg' ;
K_use : 'use' ;
K_vectored : 'vectored' ;
K_wait : 'wait' ;
K_wand : 'wand' ;
K_weak0 : 'weak0' ;
K_weak1 : 'weak1' ;
K_while : 'while' ;
K_wire : 'wire' ;
K_wor : 'wor' ;
K_xnor : 'xnor' ;
K_xor : 'xor' ;

K_fullskew	 : '$fullskew' ;
K_hold	 : '$hold' ;
K_nochange	 : '$nochange' ;
K_period	 : '$period' ;
K_recovery	 : '$recovery' ;
K_recrem	 : '$recrem' ;
K_removal	 : '$removal' ;
K_root	 : '$root' ;
K_setuphold	 : '$setuphold' ;
K_setup	 : '$setup' ;
K_skew	 : '$skew' ;
K_timeskew	 : '$timeskew' ;
K_unit	 : '$unit' ;
K_width	 : '$width' ;

// Display tasks (17.1)
K_TASK_monitoroff   : '$monitoroff' ;
K_TASK_monitoron    : '$monitoron' ;
K_TASK_monitoro     : '$monitoro' ;
K_TASK_monitorh     : '$monitorh' ;
K_TASK_monitorb     : '$monitorb' ;
K_TASK_displayo     : '$displayo' ;
K_TASK_displayh     : '$displayh' ;
K_TASK_displayb     : '$displayb' ;
K_TASK_strobeo      : '$strobeo' ;
K_TASK_strobeh      : '$strobeh' ;
K_TASK_strobeb      : '$strobeb' ;
K_TASK_monitor      : '$monitor' ;
K_TASK_display      : '$display' ;
K_TASK_writeo       : '$writeo' ;
K_TASK_writeh       : '$writeh' ;
K_TASK_writeb       : '$writeb' ;
K_TASK_strobe       : '$strobe' ;
K_TASK_write        : '$write' ;

// File I/O tasks (17.2)
K_TASK_fmonitoro    : '$fmonitoro' ;
K_TASK_fmonitorh    : '$fmonitorh' ;
K_TASK_fmonitorb    : '$fmonitorb' ;
K_TASK_fdisplayo    : '$fdisplayo' ;
K_TASK_fdisplayh    : '$fdisplayh' ;
K_TASK_fdisplayb    : '$fdisplayb' ;
K_TASK_fmonitor     : '$fmonitor' ;
K_TASK_fdisplay     : '$fdisplay' ;
K_TASK_fstrobeo     : '$fstrobeo' ;
K_TASK_fstrobeh     : '$fstrobeh' ;
K_TASK_fstrobeb     : '$fstrobeb' ;
K_TASK_sformatf     : '$sformatf' ;
K_TASK_swriteo      : '$swriteo' ;
K_TASK_swriteh      : '$swriteh' ;
K_TASK_swriteb      : '$swriteb' ;
K_TASK_sformat      : '$sformat' ;
K_TASK_fwriteo      : '$fwriteo' ;
K_TASK_fwriteh      : '$fwriteh' ;
K_TASK_fwriteb      : '$fwriteb' ;
K_TASK_fstrobe      : '$fstrobe' ;
K_TASK_ungetc       : '$ungetc' ;
K_TASK_swrite       : '$swrite' ;
K_TASK_sscanf       : '$sscanf' ;
K_TASK_rewind       : '$rewind' ;
K_TASK_fwrite       : '$fwrite' ;
K_TASK_fscanf       : '$fscanf' ;
K_TASK_fflush       : '$fflush' ;
K_TASK_ferror       : '$ferror' ;
K_TASK_fclose       : '$fclose' ;
K_TASK_ftell        : '$ftell' ;
K_TASK_fseek        : '$fseek' ;
K_TASK_fopen        : '$fopen' ;
K_TASK_fread        : '$fread' ;
K_TASK_fgets        : '$fgets' ;
K_TASK_fgetc        : '$fgetc' ;
K_TASK_feof         : '$feof' ;

// Timescale tasks (17.3)
K_TASK_printtimescale   : '$printtimescale' ;
K_TASK_timeformat       : '$timeformat' ;

// Simulation control tasks (17.4)
K_TASK_finish   : '$finish' ;
K_TASK_stop     : '$stop' ;

// PLA modeling tasks (17.5)
K_TASK_async_nand_plane : '$async$nand$plane' ;
K_TASK_async_nand_array : '$async$nand$array' ;
K_TASK_async_nor_plane  : '$async$nor$plane' ;
K_TASK_async_nor_array  : '$async$nor$array' ;
K_TASK_async_and_plane  : '$async$and$plane' ;
K_TASK_async_and_array  : '$async$and$array' ;
K_TASK_sync_nand_plane  : '$sync$nand$plane' ;
K_TASK_sync_nand_array  : '$sync$nand$array' ;
K_TASK_sync_nor_plane   : '$sync$nor$plane' ;
K_TASK_sync_nor_array   : '$sync$nor$array' ;
K_TASK_sync_and_plane   : '$sync$and$plane' ;
K_TASK_sync_and_array   : '$sync$and$array' ;
K_TASK_async_or_plane   : '$async$or$plane' ;
K_TASK_async_or_array   : '$async$or$array' ;
K_TASK_sync_or_plane    : '$sync$or$plane' ;
K_TASK_sync_or_array    : '$sync$or$array' ;

// Stochastic analysis tasks and functions (17.6)
K_FUNC_q_initialize : '$q_initialize' ;
K_FUNC_q_remove : '$q_remove' ;
K_FUNC_q_full : '$q_full' ;
K_FUNC_q_exam : '$q_exam' ;
K_FUNC_q_add : '$q_add' ;

// Simulation time functions (17.7)
K_FUNC_realtime     : '$realtime' ;
K_FUNC_stime        : '$stime' ;
K_FUNC_time         : '$time' ;

// Conversion functions (17.8)
K_FUNC_shortrealtobits  : '$shortrealtobits' ;
K_FUNC_bitstoshortreal  : '$bitstoshortreal' ;
K_FUNC_realtobits       : '$realtobits' ;
K_FUNC_bitstoreal       : '$bitstoreal' ;
K_FUNC_unsigned         : '$unsigned' ;
K_FUNC_signed           : '$signed' ;
K_FUNC_rtoi             : '$rtoi' ;
K_FUNC_itor             : '$itor' ;
K_FUNC_cast             : '$cast' ;

// Probabilistic distribution functions (17.9)
K_FUNC_random           : '$random' ;
K_FUNC_dist_exponential : '$dist_exponential' ;
K_FUNC_dist_chi_square  : '$dist_chi_square' ;
K_FUNC_dist_uniform     : '$dist_uniform' ;
K_FUNC_dist_poisson     : '$dist_poisson' ;
K_FUNC_dist_normal      : '$dist_normal' ;
K_FUNC_dist_erlang      : '$dist_erlang' ;
K_FUNC_dist_t           : '$dist_t' ;

// Command line input (17.10)
K_FUNC_test_plusargs : '$test$plusargs' ;
K_FUNC_value_plusargs  : '$value$plusargs' ;

Semicolon : ';' ;
Colon : ':' ;
Dollar : '$' ;
Dot : '.' ;
Plus : '+' ;
Minus : '-' ;
Comma : ',' ;
Question_mark : '?' ;
Exclamation_mark : '!' ;
Left_parenthes : '(' ;
Right_parenthes : ')' ;
Left_bracket : '[' ;
Right_bracket : ']' ;
Left_brace : '{' ;
Right_brace : '}' ;
Left_angle_bracket : '<' ;
Right_angle_bracket : '>' ;
Sharp : '#' ;
Underscore : '_' ;
Slash : '/' ;
// Backslash : '\\' ;
Eq : '=' ;
Eq_eq : '==' ;
Eq_eq_eq : '===' ;
Not_eq : '!=' ;
Not_eq_eq : '!==' ;
Left_angle_eq : '<=' ;
Right_angle_eq : '=>' ;
At : '@' ;
Asterisk : '*' ;
Left_arrow : '<-' ;
Right_arrow : '->' ;
Left_asterisk_arrow : '<*' ;
Right_asterisk_arrow : '*>' ;
Percent : '%' ;
AND3 : '&&&' ;
AND : '&' ;
Vertical_line : '|' ;
Apostrophe : '`' ;
Tilda : '~' ;
Hat : '^' ;
Quotation_mark : '"' ;

Ignored : Ignored_Compiler_Directives .*? New_Line -> channel(IGNORED) ;

fragment Ignored_Compiler_Directives
    // From IEEE1364-2005
    : ('`begin_keywords'
    | '`celldefine'
    | '`default_nettype'
    | '`end_keywords'
    | '`endcelldefine'
    | '`nounconnected_drive'
    | '`pragma'
    | '`resetall'
    | '`unconnected_drive')
    ;

// LINE : '`line' {  } -> channel(HIDDEN_CHANNEL), mode(LineMode);
LINE : '`line' {  } -> channel(LINE_CHANNEL), mode(LineMode);

Real_number
   : Unsigned_number Dot Unsigned_number | Unsigned_number (Dot Unsigned_number)? [eE] ([+-])? Unsigned_number
   ;

Decimal_number
   : Unsigned_number | (Size)? Decimal_base Unsigned_number | (Size)? Decimal_base X_digit (Underscore)* | (Size)? Decimal_base Z_digit (Underscore)*
   ;

Binary_number
   : (Size)? Binary_base Binary_value
   ;

Octal_number
   : (Size)? Octal_base Octal_value
   ;

Hex_number
   : (Size)? Hex_base Hex_value
   ;

fragment Escaped_New_Line 
   : '\\' '\r'? '\n' 
   ;

fragment New_Line 
   : '\r'? '\n' 
   ;

fragment Sign
   : [+-]
   ;

fragment Size
   : Non_zero_unsigned_number
   ;

fragment Non_zero_unsigned_number
   : Non_zero_decimal_digit (Underscore | Decimal_digit)*
   ;

fragment Unsigned_number
   : Decimal_digit (Underscore | Decimal_digit)*
   ;

fragment Binary_value
   : Binary_digit (Underscore | Binary_digit)*
   ;

fragment Octal_value
   : Octal_digit (Underscore | Octal_digit)*
   ;

fragment Hex_value
   : Hex_digit (Underscore | Hex_digit)*
   ;

fragment Decimal_base
   : '\'' [sS]? [dD]
   ;

fragment Binary_base
   : '\'' [sS]? [bB]
   ;

fragment Octal_base
   : '\'' [sS]? [oO]
   ;

fragment Hex_base
   : '\'' [sS]? [hH]
   ;

fragment Non_zero_decimal_digit
   : [1-9]
   ;

fragment Decimal_digit
   : [0-9]
   ;

fragment Binary_digit
   : X_digit | Z_digit | [01]
   ;

fragment Octal_digit
   : X_digit | Z_digit | [0-7]
   ;

fragment Hex_digit
   : X_digit | Z_digit | [0-9a-fA-F]
   ;

fragment X_digit
   : [xX]
   ;

fragment Z_digit
   : [zZ?]
   ;

Single_Bit_Zero : '1\'b0' | '1\'B0' ;

Single_Bit_One : '1\'b1' | '1\'B1' ;

Single_Bit_X : '1\'bx' | '1\'BX' ;

All_Binary_Zero : '\'b0' | '\'B0' ;

All_Binary_One : '\'b1' | '\'B1' ;

Escaped_identifier
   : '\\' ('\u0021'..'\u007E')+ ~ [ \r\t\n]*
   ;

Simple_identifier
   : [a-zA-Z_] [a-zA-Z0-9_$]*
   ;

Dollar_identifier
   : Dollar [a-zA-Z0-9_$] [a-zA-Z0-9_$]*
   ;

Time_identifier
   : Unsigned_number [ /t]* [mnpf]? 's'
   ;

String_const
   : '"' (~ [\n\r])* '"'
   ;

One_line_comment
   : '//' .*? '\r'? '\n' -> channel (HIDDEN)
   ;

Block_comment
   : '/*' .*? '*/' -> channel (HIDDEN)
   ;

Zero : '0' ;

One : '1' ;


Edge_descriptor
    : Rising_Edge
    | Falling_Edge
    | X_Rising_Edge
    | X_Falling_Edge
    ;

Rising_Edge : '01' ;

Falling_Edge : '10' ;

X_Rising_Edge : 'x0' | 'x1' | 'X0' | 'X1' | 'z0' | 'z1' | 'Z0' | 'Z1' ;

X_Falling_Edge : '0x' | '1x' | '0X' | '1X' | '0z' | '1z' | '0Z' | '1Z' ;

White_space
   : (' ' | '\t' | Escaped_New_Line | New_Line)+ -> channel(HIDDEN)
   ;

Bad_character
   : .
   ;

mode LineMode;

LINE_MODE_NUMBER: (Decimal_digit)+ 
{
   CurrentLineNumber, _ = strconv.Atoi(l.GetText())
}
-> channel(LINE_CHANNEL);

LINE_MODE_STRING: String_const
{
   CurrentFileName = l.GetText();
   CurrentFileName = strings.Replace(CurrentFileName, "\"", "", -1)
   // fmt.Printf("current filename: %s\n", CurrentFileName)
}
-> channel(LINE_CHANNEL);

LINE_MODE_NEW_LINE: '\r'? '\n' -> channel(HIDDEN_CHANNEL), mode(DEFAULT_MODE) ;

LINE_MODE_WHITE_SPACE: [ \t]+ -> channel(HIDDEN_CHANNEL) ;

LINE_MODE: ~["0-9\r\n]+ -> channel(HIDDEN_CHANNEL) ;


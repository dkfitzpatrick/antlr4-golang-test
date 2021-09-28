// Code generated from VerilogParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package verilog // VerilogParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVerilogParserListener is a complete listener for a parse tree produced by VerilogParser.
type BaseVerilogParserListener struct{}

var _ VerilogParserListener = &BaseVerilogParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVerilogParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVerilogParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVerilogParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVerilogParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseVerilogParserListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseVerilogParserListener) ExitDirective(ctx *DirectiveContext) {}

// EnterTimescale_directive is called when production timescale_directive is entered.
func (s *BaseVerilogParserListener) EnterTimescale_directive(ctx *Timescale_directiveContext) {}

// ExitTimescale_directive is called when production timescale_directive is exited.
func (s *BaseVerilogParserListener) ExitTimescale_directive(ctx *Timescale_directiveContext) {}

// EnterConfig_declaration is called when production config_declaration is entered.
func (s *BaseVerilogParserListener) EnterConfig_declaration(ctx *Config_declarationContext) {}

// ExitConfig_declaration is called when production config_declaration is exited.
func (s *BaseVerilogParserListener) ExitConfig_declaration(ctx *Config_declarationContext) {}

// EnterDesign_statement is called when production design_statement is entered.
func (s *BaseVerilogParserListener) EnterDesign_statement(ctx *Design_statementContext) {}

// ExitDesign_statement is called when production design_statement is exited.
func (s *BaseVerilogParserListener) ExitDesign_statement(ctx *Design_statementContext) {}

// EnterConfig_rule_statement is called when production config_rule_statement is entered.
func (s *BaseVerilogParserListener) EnterConfig_rule_statement(ctx *Config_rule_statementContext) {}

// ExitConfig_rule_statement is called when production config_rule_statement is exited.
func (s *BaseVerilogParserListener) ExitConfig_rule_statement(ctx *Config_rule_statementContext) {}

// EnterDefault_clause is called when production default_clause is entered.
func (s *BaseVerilogParserListener) EnterDefault_clause(ctx *Default_clauseContext) {}

// ExitDefault_clause is called when production default_clause is exited.
func (s *BaseVerilogParserListener) ExitDefault_clause(ctx *Default_clauseContext) {}

// EnterInst_clause is called when production inst_clause is entered.
func (s *BaseVerilogParserListener) EnterInst_clause(ctx *Inst_clauseContext) {}

// ExitInst_clause is called when production inst_clause is exited.
func (s *BaseVerilogParserListener) ExitInst_clause(ctx *Inst_clauseContext) {}

// EnterInst_name is called when production inst_name is entered.
func (s *BaseVerilogParserListener) EnterInst_name(ctx *Inst_nameContext) {}

// ExitInst_name is called when production inst_name is exited.
func (s *BaseVerilogParserListener) ExitInst_name(ctx *Inst_nameContext) {}

// EnterLiblist_clause is called when production liblist_clause is entered.
func (s *BaseVerilogParserListener) EnterLiblist_clause(ctx *Liblist_clauseContext) {}

// ExitLiblist_clause is called when production liblist_clause is exited.
func (s *BaseVerilogParserListener) ExitLiblist_clause(ctx *Liblist_clauseContext) {}

// EnterCell_clause is called when production cell_clause is entered.
func (s *BaseVerilogParserListener) EnterCell_clause(ctx *Cell_clauseContext) {}

// ExitCell_clause is called when production cell_clause is exited.
func (s *BaseVerilogParserListener) ExitCell_clause(ctx *Cell_clauseContext) {}

// EnterUse_clause is called when production use_clause is entered.
func (s *BaseVerilogParserListener) EnterUse_clause(ctx *Use_clauseContext) {}

// ExitUse_clause is called when production use_clause is exited.
func (s *BaseVerilogParserListener) ExitUse_clause(ctx *Use_clauseContext) {}

// EnterSource_text is called when production source_text is entered.
func (s *BaseVerilogParserListener) EnterSource_text(ctx *Source_textContext) {}

// ExitSource_text is called when production source_text is exited.
func (s *BaseVerilogParserListener) ExitSource_text(ctx *Source_textContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseVerilogParserListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseVerilogParserListener) ExitDescription(ctx *DescriptionContext) {}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *BaseVerilogParserListener) EnterModule_declaration(ctx *Module_declarationContext) {}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *BaseVerilogParserListener) ExitModule_declaration(ctx *Module_declarationContext) {}

// EnterModule_keyword is called when production module_keyword is entered.
func (s *BaseVerilogParserListener) EnterModule_keyword(ctx *Module_keywordContext) {}

// ExitModule_keyword is called when production module_keyword is exited.
func (s *BaseVerilogParserListener) ExitModule_keyword(ctx *Module_keywordContext) {}

// EnterModule_parameter_port_list is called when production module_parameter_port_list is entered.
func (s *BaseVerilogParserListener) EnterModule_parameter_port_list(ctx *Module_parameter_port_listContext) {
}

// ExitModule_parameter_port_list is called when production module_parameter_port_list is exited.
func (s *BaseVerilogParserListener) ExitModule_parameter_port_list(ctx *Module_parameter_port_listContext) {
}

// EnterList_of_ports is called when production list_of_ports is entered.
func (s *BaseVerilogParserListener) EnterList_of_ports(ctx *List_of_portsContext) {}

// ExitList_of_ports is called when production list_of_ports is exited.
func (s *BaseVerilogParserListener) ExitList_of_ports(ctx *List_of_portsContext) {}

// EnterList_of_port_declarations is called when production list_of_port_declarations is entered.
func (s *BaseVerilogParserListener) EnterList_of_port_declarations(ctx *List_of_port_declarationsContext) {
}

// ExitList_of_port_declarations is called when production list_of_port_declarations is exited.
func (s *BaseVerilogParserListener) ExitList_of_port_declarations(ctx *List_of_port_declarationsContext) {
}

// EnterPort is called when production port is entered.
func (s *BaseVerilogParserListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseVerilogParserListener) ExitPort(ctx *PortContext) {}

// EnterPort_expression is called when production port_expression is entered.
func (s *BaseVerilogParserListener) EnterPort_expression(ctx *Port_expressionContext) {}

// ExitPort_expression is called when production port_expression is exited.
func (s *BaseVerilogParserListener) ExitPort_expression(ctx *Port_expressionContext) {}

// EnterPort_reference is called when production port_reference is entered.
func (s *BaseVerilogParserListener) EnterPort_reference(ctx *Port_referenceContext) {}

// ExitPort_reference is called when production port_reference is exited.
func (s *BaseVerilogParserListener) ExitPort_reference(ctx *Port_referenceContext) {}

// EnterPort_declaration is called when production port_declaration is entered.
func (s *BaseVerilogParserListener) EnterPort_declaration(ctx *Port_declarationContext) {}

// ExitPort_declaration is called when production port_declaration is exited.
func (s *BaseVerilogParserListener) ExitPort_declaration(ctx *Port_declarationContext) {}

// EnterModule_item is called when production module_item is entered.
func (s *BaseVerilogParserListener) EnterModule_item(ctx *Module_itemContext) {}

// ExitModule_item is called when production module_item is exited.
func (s *BaseVerilogParserListener) ExitModule_item(ctx *Module_itemContext) {}

// EnterModule_or_generate_item is called when production module_or_generate_item is entered.
func (s *BaseVerilogParserListener) EnterModule_or_generate_item(ctx *Module_or_generate_itemContext) {
}

// ExitModule_or_generate_item is called when production module_or_generate_item is exited.
func (s *BaseVerilogParserListener) ExitModule_or_generate_item(ctx *Module_or_generate_itemContext) {
}

// EnterModule_or_generate_item_declaration is called when production module_or_generate_item_declaration is entered.
func (s *BaseVerilogParserListener) EnterModule_or_generate_item_declaration(ctx *Module_or_generate_item_declarationContext) {
}

// ExitModule_or_generate_item_declaration is called when production module_or_generate_item_declaration is exited.
func (s *BaseVerilogParserListener) ExitModule_or_generate_item_declaration(ctx *Module_or_generate_item_declarationContext) {
}

// EnterParameter_override is called when production parameter_override is entered.
func (s *BaseVerilogParserListener) EnterParameter_override(ctx *Parameter_overrideContext) {}

// ExitParameter_override is called when production parameter_override is exited.
func (s *BaseVerilogParserListener) ExitParameter_override(ctx *Parameter_overrideContext) {}

// EnterLocal_parameter_declaration is called when production local_parameter_declaration is entered.
func (s *BaseVerilogParserListener) EnterLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// ExitLocal_parameter_declaration is called when production local_parameter_declaration is exited.
func (s *BaseVerilogParserListener) ExitLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseVerilogParserListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseVerilogParserListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterParameter_declaration_ is called when production parameter_declaration_ is entered.
func (s *BaseVerilogParserListener) EnterParameter_declaration_(ctx *Parameter_declaration_Context) {}

// ExitParameter_declaration_ is called when production parameter_declaration_ is exited.
func (s *BaseVerilogParserListener) ExitParameter_declaration_(ctx *Parameter_declaration_Context) {}

// EnterSpecparam_declaration is called when production specparam_declaration is entered.
func (s *BaseVerilogParserListener) EnterSpecparam_declaration(ctx *Specparam_declarationContext) {}

// ExitSpecparam_declaration is called when production specparam_declaration is exited.
func (s *BaseVerilogParserListener) ExitSpecparam_declaration(ctx *Specparam_declarationContext) {}

// EnterInout_declaration is called when production inout_declaration is entered.
func (s *BaseVerilogParserListener) EnterInout_declaration(ctx *Inout_declarationContext) {}

// ExitInout_declaration is called when production inout_declaration is exited.
func (s *BaseVerilogParserListener) ExitInout_declaration(ctx *Inout_declarationContext) {}

// EnterInput_declaration is called when production input_declaration is entered.
func (s *BaseVerilogParserListener) EnterInput_declaration(ctx *Input_declarationContext) {}

// ExitInput_declaration is called when production input_declaration is exited.
func (s *BaseVerilogParserListener) ExitInput_declaration(ctx *Input_declarationContext) {}

// EnterOutput_declaration is called when production output_declaration is entered.
func (s *BaseVerilogParserListener) EnterOutput_declaration(ctx *Output_declarationContext) {}

// ExitOutput_declaration is called when production output_declaration is exited.
func (s *BaseVerilogParserListener) ExitOutput_declaration(ctx *Output_declarationContext) {}

// EnterEvent_declaration is called when production event_declaration is entered.
func (s *BaseVerilogParserListener) EnterEvent_declaration(ctx *Event_declarationContext) {}

// ExitEvent_declaration is called when production event_declaration is exited.
func (s *BaseVerilogParserListener) ExitEvent_declaration(ctx *Event_declarationContext) {}

// EnterGenvar_declaration is called when production genvar_declaration is entered.
func (s *BaseVerilogParserListener) EnterGenvar_declaration(ctx *Genvar_declarationContext) {}

// ExitGenvar_declaration is called when production genvar_declaration is exited.
func (s *BaseVerilogParserListener) ExitGenvar_declaration(ctx *Genvar_declarationContext) {}

// EnterInteger_declaration is called when production integer_declaration is entered.
func (s *BaseVerilogParserListener) EnterInteger_declaration(ctx *Integer_declarationContext) {}

// ExitInteger_declaration is called when production integer_declaration is exited.
func (s *BaseVerilogParserListener) ExitInteger_declaration(ctx *Integer_declarationContext) {}

// EnterTime_declaration is called when production time_declaration is entered.
func (s *BaseVerilogParserListener) EnterTime_declaration(ctx *Time_declarationContext) {}

// ExitTime_declaration is called when production time_declaration is exited.
func (s *BaseVerilogParserListener) ExitTime_declaration(ctx *Time_declarationContext) {}

// EnterReal_declaration is called when production real_declaration is entered.
func (s *BaseVerilogParserListener) EnterReal_declaration(ctx *Real_declarationContext) {}

// ExitReal_declaration is called when production real_declaration is exited.
func (s *BaseVerilogParserListener) ExitReal_declaration(ctx *Real_declarationContext) {}

// EnterRealtime_declaration is called when production realtime_declaration is entered.
func (s *BaseVerilogParserListener) EnterRealtime_declaration(ctx *Realtime_declarationContext) {}

// ExitRealtime_declaration is called when production realtime_declaration is exited.
func (s *BaseVerilogParserListener) ExitRealtime_declaration(ctx *Realtime_declarationContext) {}

// EnterReg_declaration is called when production reg_declaration is entered.
func (s *BaseVerilogParserListener) EnterReg_declaration(ctx *Reg_declarationContext) {}

// ExitReg_declaration is called when production reg_declaration is exited.
func (s *BaseVerilogParserListener) ExitReg_declaration(ctx *Reg_declarationContext) {}

// EnterNet_declaration is called when production net_declaration is entered.
func (s *BaseVerilogParserListener) EnterNet_declaration(ctx *Net_declarationContext) {}

// ExitNet_declaration is called when production net_declaration is exited.
func (s *BaseVerilogParserListener) ExitNet_declaration(ctx *Net_declarationContext) {}

// EnterNet_type is called when production net_type is entered.
func (s *BaseVerilogParserListener) EnterNet_type(ctx *Net_typeContext) {}

// ExitNet_type is called when production net_type is exited.
func (s *BaseVerilogParserListener) ExitNet_type(ctx *Net_typeContext) {}

// EnterOutput_variable_type is called when production output_variable_type is entered.
func (s *BaseVerilogParserListener) EnterOutput_variable_type(ctx *Output_variable_typeContext) {}

// ExitOutput_variable_type is called when production output_variable_type is exited.
func (s *BaseVerilogParserListener) ExitOutput_variable_type(ctx *Output_variable_typeContext) {}

// EnterReal_type is called when production real_type is entered.
func (s *BaseVerilogParserListener) EnterReal_type(ctx *Real_typeContext) {}

// ExitReal_type is called when production real_type is exited.
func (s *BaseVerilogParserListener) ExitReal_type(ctx *Real_typeContext) {}

// EnterVariable_type is called when production variable_type is entered.
func (s *BaseVerilogParserListener) EnterVariable_type(ctx *Variable_typeContext) {}

// ExitVariable_type is called when production variable_type is exited.
func (s *BaseVerilogParserListener) ExitVariable_type(ctx *Variable_typeContext) {}

// EnterDrive_strength is called when production drive_strength is entered.
func (s *BaseVerilogParserListener) EnterDrive_strength(ctx *Drive_strengthContext) {}

// ExitDrive_strength is called when production drive_strength is exited.
func (s *BaseVerilogParserListener) ExitDrive_strength(ctx *Drive_strengthContext) {}

// EnterStrength0 is called when production strength0 is entered.
func (s *BaseVerilogParserListener) EnterStrength0(ctx *Strength0Context) {}

// ExitStrength0 is called when production strength0 is exited.
func (s *BaseVerilogParserListener) ExitStrength0(ctx *Strength0Context) {}

// EnterStrength1 is called when production strength1 is entered.
func (s *BaseVerilogParserListener) EnterStrength1(ctx *Strength1Context) {}

// ExitStrength1 is called when production strength1 is exited.
func (s *BaseVerilogParserListener) ExitStrength1(ctx *Strength1Context) {}

// EnterCharge_strength is called when production charge_strength is entered.
func (s *BaseVerilogParserListener) EnterCharge_strength(ctx *Charge_strengthContext) {}

// ExitCharge_strength is called when production charge_strength is exited.
func (s *BaseVerilogParserListener) ExitCharge_strength(ctx *Charge_strengthContext) {}

// EnterDelay3 is called when production delay3 is entered.
func (s *BaseVerilogParserListener) EnterDelay3(ctx *Delay3Context) {}

// ExitDelay3 is called when production delay3 is exited.
func (s *BaseVerilogParserListener) ExitDelay3(ctx *Delay3Context) {}

// EnterDelay2 is called when production delay2 is entered.
func (s *BaseVerilogParserListener) EnterDelay2(ctx *Delay2Context) {}

// ExitDelay2 is called when production delay2 is exited.
func (s *BaseVerilogParserListener) ExitDelay2(ctx *Delay2Context) {}

// EnterDelay_value is called when production delay_value is entered.
func (s *BaseVerilogParserListener) EnterDelay_value(ctx *Delay_valueContext) {}

// ExitDelay_value is called when production delay_value is exited.
func (s *BaseVerilogParserListener) ExitDelay_value(ctx *Delay_valueContext) {}

// EnterList_of_event_identifiers is called when production list_of_event_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_event_identifiers(ctx *List_of_event_identifiersContext) {
}

// ExitList_of_event_identifiers is called when production list_of_event_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_event_identifiers(ctx *List_of_event_identifiersContext) {
}

// EnterList_of_net_identifiers is called when production list_of_net_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_net_identifiers(ctx *List_of_net_identifiersContext) {
}

// ExitList_of_net_identifiers is called when production list_of_net_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_net_identifiers(ctx *List_of_net_identifiersContext) {
}

// EnterList_of_genvar_identifiers is called when production list_of_genvar_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_genvar_identifiers(ctx *List_of_genvar_identifiersContext) {
}

// ExitList_of_genvar_identifiers is called when production list_of_genvar_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_genvar_identifiers(ctx *List_of_genvar_identifiersContext) {
}

// EnterList_of_port_identifiers is called when production list_of_port_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_port_identifiers(ctx *List_of_port_identifiersContext) {
}

// ExitList_of_port_identifiers is called when production list_of_port_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_port_identifiers(ctx *List_of_port_identifiersContext) {
}

// EnterList_of_net_decl_assignments is called when production list_of_net_decl_assignments is entered.
func (s *BaseVerilogParserListener) EnterList_of_net_decl_assignments(ctx *List_of_net_decl_assignmentsContext) {
}

// ExitList_of_net_decl_assignments is called when production list_of_net_decl_assignments is exited.
func (s *BaseVerilogParserListener) ExitList_of_net_decl_assignments(ctx *List_of_net_decl_assignmentsContext) {
}

// EnterList_of_param_assignments is called when production list_of_param_assignments is entered.
func (s *BaseVerilogParserListener) EnterList_of_param_assignments(ctx *List_of_param_assignmentsContext) {
}

// ExitList_of_param_assignments is called when production list_of_param_assignments is exited.
func (s *BaseVerilogParserListener) ExitList_of_param_assignments(ctx *List_of_param_assignmentsContext) {
}

// EnterList_of_specparam_assignments is called when production list_of_specparam_assignments is entered.
func (s *BaseVerilogParserListener) EnterList_of_specparam_assignments(ctx *List_of_specparam_assignmentsContext) {
}

// ExitList_of_specparam_assignments is called when production list_of_specparam_assignments is exited.
func (s *BaseVerilogParserListener) ExitList_of_specparam_assignments(ctx *List_of_specparam_assignmentsContext) {
}

// EnterList_of_real_identifiers is called when production list_of_real_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_real_identifiers(ctx *List_of_real_identifiersContext) {
}

// ExitList_of_real_identifiers is called when production list_of_real_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_real_identifiers(ctx *List_of_real_identifiersContext) {
}

// EnterList_of_variable_identifiers is called when production list_of_variable_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_variable_identifiers(ctx *List_of_variable_identifiersContext) {
}

// ExitList_of_variable_identifiers is called when production list_of_variable_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_variable_identifiers(ctx *List_of_variable_identifiersContext) {
}

// EnterList_of_variable_port_identifiers is called when production list_of_variable_port_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_variable_port_identifiers(ctx *List_of_variable_port_identifiersContext) {
}

// ExitList_of_variable_port_identifiers is called when production list_of_variable_port_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_variable_port_identifiers(ctx *List_of_variable_port_identifiersContext) {
}

// EnterNet_decl_assignment is called when production net_decl_assignment is entered.
func (s *BaseVerilogParserListener) EnterNet_decl_assignment(ctx *Net_decl_assignmentContext) {}

// ExitNet_decl_assignment is called when production net_decl_assignment is exited.
func (s *BaseVerilogParserListener) ExitNet_decl_assignment(ctx *Net_decl_assignmentContext) {}

// EnterParam_assignment is called when production param_assignment is entered.
func (s *BaseVerilogParserListener) EnterParam_assignment(ctx *Param_assignmentContext) {}

// ExitParam_assignment is called when production param_assignment is exited.
func (s *BaseVerilogParserListener) ExitParam_assignment(ctx *Param_assignmentContext) {}

// EnterSpecparam_assignment is called when production specparam_assignment is entered.
func (s *BaseVerilogParserListener) EnterSpecparam_assignment(ctx *Specparam_assignmentContext) {}

// ExitSpecparam_assignment is called when production specparam_assignment is exited.
func (s *BaseVerilogParserListener) ExitSpecparam_assignment(ctx *Specparam_assignmentContext) {}

// EnterPulse_control_specparam is called when production pulse_control_specparam is entered.
func (s *BaseVerilogParserListener) EnterPulse_control_specparam(ctx *Pulse_control_specparamContext) {
}

// ExitPulse_control_specparam is called when production pulse_control_specparam is exited.
func (s *BaseVerilogParserListener) ExitPulse_control_specparam(ctx *Pulse_control_specparamContext) {
}

// EnterError_limit_value is called when production error_limit_value is entered.
func (s *BaseVerilogParserListener) EnterError_limit_value(ctx *Error_limit_valueContext) {}

// ExitError_limit_value is called when production error_limit_value is exited.
func (s *BaseVerilogParserListener) ExitError_limit_value(ctx *Error_limit_valueContext) {}

// EnterReject_limit_value is called when production reject_limit_value is entered.
func (s *BaseVerilogParserListener) EnterReject_limit_value(ctx *Reject_limit_valueContext) {}

// ExitReject_limit_value is called when production reject_limit_value is exited.
func (s *BaseVerilogParserListener) ExitReject_limit_value(ctx *Reject_limit_valueContext) {}

// EnterLimit_value is called when production limit_value is entered.
func (s *BaseVerilogParserListener) EnterLimit_value(ctx *Limit_valueContext) {}

// ExitLimit_value is called when production limit_value is exited.
func (s *BaseVerilogParserListener) ExitLimit_value(ctx *Limit_valueContext) {}

// EnterDimension is called when production dimension is entered.
func (s *BaseVerilogParserListener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *BaseVerilogParserListener) ExitDimension(ctx *DimensionContext) {}

// EnterRange_ is called when production range_ is entered.
func (s *BaseVerilogParserListener) EnterRange_(ctx *Range_Context) {}

// ExitRange_ is called when production range_ is exited.
func (s *BaseVerilogParserListener) ExitRange_(ctx *Range_Context) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseVerilogParserListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseVerilogParserListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterFunction_item_declaration is called when production function_item_declaration is entered.
func (s *BaseVerilogParserListener) EnterFunction_item_declaration(ctx *Function_item_declarationContext) {
}

// ExitFunction_item_declaration is called when production function_item_declaration is exited.
func (s *BaseVerilogParserListener) ExitFunction_item_declaration(ctx *Function_item_declarationContext) {
}

// EnterFunction_port_list is called when production function_port_list is entered.
func (s *BaseVerilogParserListener) EnterFunction_port_list(ctx *Function_port_listContext) {}

// ExitFunction_port_list is called when production function_port_list is exited.
func (s *BaseVerilogParserListener) ExitFunction_port_list(ctx *Function_port_listContext) {}

// EnterFunction_port is called when production function_port is entered.
func (s *BaseVerilogParserListener) EnterFunction_port(ctx *Function_portContext) {}

// ExitFunction_port is called when production function_port is exited.
func (s *BaseVerilogParserListener) ExitFunction_port(ctx *Function_portContext) {}

// EnterRange_or_type is called when production range_or_type is entered.
func (s *BaseVerilogParserListener) EnterRange_or_type(ctx *Range_or_typeContext) {}

// ExitRange_or_type is called when production range_or_type is exited.
func (s *BaseVerilogParserListener) ExitRange_or_type(ctx *Range_or_typeContext) {}

// EnterTask_declaration is called when production task_declaration is entered.
func (s *BaseVerilogParserListener) EnterTask_declaration(ctx *Task_declarationContext) {}

// ExitTask_declaration is called when production task_declaration is exited.
func (s *BaseVerilogParserListener) ExitTask_declaration(ctx *Task_declarationContext) {}

// EnterTask_item_declaration is called when production task_item_declaration is entered.
func (s *BaseVerilogParserListener) EnterTask_item_declaration(ctx *Task_item_declarationContext) {}

// ExitTask_item_declaration is called when production task_item_declaration is exited.
func (s *BaseVerilogParserListener) ExitTask_item_declaration(ctx *Task_item_declarationContext) {}

// EnterTask_port_list is called when production task_port_list is entered.
func (s *BaseVerilogParserListener) EnterTask_port_list(ctx *Task_port_listContext) {}

// ExitTask_port_list is called when production task_port_list is exited.
func (s *BaseVerilogParserListener) ExitTask_port_list(ctx *Task_port_listContext) {}

// EnterTask_port_item is called when production task_port_item is entered.
func (s *BaseVerilogParserListener) EnterTask_port_item(ctx *Task_port_itemContext) {}

// ExitTask_port_item is called when production task_port_item is exited.
func (s *BaseVerilogParserListener) ExitTask_port_item(ctx *Task_port_itemContext) {}

// EnterTf_decl_header is called when production tf_decl_header is entered.
func (s *BaseVerilogParserListener) EnterTf_decl_header(ctx *Tf_decl_headerContext) {}

// ExitTf_decl_header is called when production tf_decl_header is exited.
func (s *BaseVerilogParserListener) ExitTf_decl_header(ctx *Tf_decl_headerContext) {}

// EnterTf_declaration is called when production tf_declaration is entered.
func (s *BaseVerilogParserListener) EnterTf_declaration(ctx *Tf_declarationContext) {}

// ExitTf_declaration is called when production tf_declaration is exited.
func (s *BaseVerilogParserListener) ExitTf_declaration(ctx *Tf_declarationContext) {}

// EnterTask_port_type is called when production task_port_type is entered.
func (s *BaseVerilogParserListener) EnterTask_port_type(ctx *Task_port_typeContext) {}

// ExitTask_port_type is called when production task_port_type is exited.
func (s *BaseVerilogParserListener) ExitTask_port_type(ctx *Task_port_typeContext) {}

// EnterBlock_item_declaration is called when production block_item_declaration is entered.
func (s *BaseVerilogParserListener) EnterBlock_item_declaration(ctx *Block_item_declarationContext) {}

// ExitBlock_item_declaration is called when production block_item_declaration is exited.
func (s *BaseVerilogParserListener) ExitBlock_item_declaration(ctx *Block_item_declarationContext) {}

// EnterBlock_reg_declaration is called when production block_reg_declaration is entered.
func (s *BaseVerilogParserListener) EnterBlock_reg_declaration(ctx *Block_reg_declarationContext) {}

// ExitBlock_reg_declaration is called when production block_reg_declaration is exited.
func (s *BaseVerilogParserListener) ExitBlock_reg_declaration(ctx *Block_reg_declarationContext) {}

// EnterList_of_block_variable_identifiers is called when production list_of_block_variable_identifiers is entered.
func (s *BaseVerilogParserListener) EnterList_of_block_variable_identifiers(ctx *List_of_block_variable_identifiersContext) {
}

// ExitList_of_block_variable_identifiers is called when production list_of_block_variable_identifiers is exited.
func (s *BaseVerilogParserListener) ExitList_of_block_variable_identifiers(ctx *List_of_block_variable_identifiersContext) {
}

// EnterBlock_variable_type is called when production block_variable_type is entered.
func (s *BaseVerilogParserListener) EnterBlock_variable_type(ctx *Block_variable_typeContext) {}

// ExitBlock_variable_type is called when production block_variable_type is exited.
func (s *BaseVerilogParserListener) ExitBlock_variable_type(ctx *Block_variable_typeContext) {}

// EnterGate_instantiation is called when production gate_instantiation is entered.
func (s *BaseVerilogParserListener) EnterGate_instantiation(ctx *Gate_instantiationContext) {}

// ExitGate_instantiation is called when production gate_instantiation is exited.
func (s *BaseVerilogParserListener) ExitGate_instantiation(ctx *Gate_instantiationContext) {}

// EnterCmos_switch_instance is called when production cmos_switch_instance is entered.
func (s *BaseVerilogParserListener) EnterCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// ExitCmos_switch_instance is called when production cmos_switch_instance is exited.
func (s *BaseVerilogParserListener) ExitCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// EnterEnable_gate_instance is called when production enable_gate_instance is entered.
func (s *BaseVerilogParserListener) EnterEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// ExitEnable_gate_instance is called when production enable_gate_instance is exited.
func (s *BaseVerilogParserListener) ExitEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// EnterMos_switch_instance is called when production mos_switch_instance is entered.
func (s *BaseVerilogParserListener) EnterMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// ExitMos_switch_instance is called when production mos_switch_instance is exited.
func (s *BaseVerilogParserListener) ExitMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// EnterN_input_gate_instance is called when production n_input_gate_instance is entered.
func (s *BaseVerilogParserListener) EnterN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// ExitN_input_gate_instance is called when production n_input_gate_instance is exited.
func (s *BaseVerilogParserListener) ExitN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// EnterN_output_gate_instance is called when production n_output_gate_instance is entered.
func (s *BaseVerilogParserListener) EnterN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// ExitN_output_gate_instance is called when production n_output_gate_instance is exited.
func (s *BaseVerilogParserListener) ExitN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// EnterPass_switch_instance is called when production pass_switch_instance is entered.
func (s *BaseVerilogParserListener) EnterPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// ExitPass_switch_instance is called when production pass_switch_instance is exited.
func (s *BaseVerilogParserListener) ExitPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// EnterPass_enable_switch_instance is called when production pass_enable_switch_instance is entered.
func (s *BaseVerilogParserListener) EnterPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// ExitPass_enable_switch_instance is called when production pass_enable_switch_instance is exited.
func (s *BaseVerilogParserListener) ExitPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// EnterPull_gate_instance is called when production pull_gate_instance is entered.
func (s *BaseVerilogParserListener) EnterPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// ExitPull_gate_instance is called when production pull_gate_instance is exited.
func (s *BaseVerilogParserListener) ExitPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// EnterName_of_gate_instance is called when production name_of_gate_instance is entered.
func (s *BaseVerilogParserListener) EnterName_of_gate_instance(ctx *Name_of_gate_instanceContext) {}

// ExitName_of_gate_instance is called when production name_of_gate_instance is exited.
func (s *BaseVerilogParserListener) ExitName_of_gate_instance(ctx *Name_of_gate_instanceContext) {}

// EnterPulldown_strength is called when production pulldown_strength is entered.
func (s *BaseVerilogParserListener) EnterPulldown_strength(ctx *Pulldown_strengthContext) {}

// ExitPulldown_strength is called when production pulldown_strength is exited.
func (s *BaseVerilogParserListener) ExitPulldown_strength(ctx *Pulldown_strengthContext) {}

// EnterPullup_strength is called when production pullup_strength is entered.
func (s *BaseVerilogParserListener) EnterPullup_strength(ctx *Pullup_strengthContext) {}

// ExitPullup_strength is called when production pullup_strength is exited.
func (s *BaseVerilogParserListener) ExitPullup_strength(ctx *Pullup_strengthContext) {}

// EnterEnable_terminal is called when production enable_terminal is entered.
func (s *BaseVerilogParserListener) EnterEnable_terminal(ctx *Enable_terminalContext) {}

// ExitEnable_terminal is called when production enable_terminal is exited.
func (s *BaseVerilogParserListener) ExitEnable_terminal(ctx *Enable_terminalContext) {}

// EnterNcontrol_terminal is called when production ncontrol_terminal is entered.
func (s *BaseVerilogParserListener) EnterNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// ExitNcontrol_terminal is called when production ncontrol_terminal is exited.
func (s *BaseVerilogParserListener) ExitNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// EnterPcontrol_terminal is called when production pcontrol_terminal is entered.
func (s *BaseVerilogParserListener) EnterPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// ExitPcontrol_terminal is called when production pcontrol_terminal is exited.
func (s *BaseVerilogParserListener) ExitPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// EnterInput_terminal is called when production input_terminal is entered.
func (s *BaseVerilogParserListener) EnterInput_terminal(ctx *Input_terminalContext) {}

// ExitInput_terminal is called when production input_terminal is exited.
func (s *BaseVerilogParserListener) ExitInput_terminal(ctx *Input_terminalContext) {}

// EnterInout_terminal is called when production inout_terminal is entered.
func (s *BaseVerilogParserListener) EnterInout_terminal(ctx *Inout_terminalContext) {}

// ExitInout_terminal is called when production inout_terminal is exited.
func (s *BaseVerilogParserListener) ExitInout_terminal(ctx *Inout_terminalContext) {}

// EnterOutput_terminal is called when production output_terminal is entered.
func (s *BaseVerilogParserListener) EnterOutput_terminal(ctx *Output_terminalContext) {}

// ExitOutput_terminal is called when production output_terminal is exited.
func (s *BaseVerilogParserListener) ExitOutput_terminal(ctx *Output_terminalContext) {}

// EnterCmos_switchtype is called when production cmos_switchtype is entered.
func (s *BaseVerilogParserListener) EnterCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// ExitCmos_switchtype is called when production cmos_switchtype is exited.
func (s *BaseVerilogParserListener) ExitCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// EnterEnable_gatetype is called when production enable_gatetype is entered.
func (s *BaseVerilogParserListener) EnterEnable_gatetype(ctx *Enable_gatetypeContext) {}

// ExitEnable_gatetype is called when production enable_gatetype is exited.
func (s *BaseVerilogParserListener) ExitEnable_gatetype(ctx *Enable_gatetypeContext) {}

// EnterMos_switchtype is called when production mos_switchtype is entered.
func (s *BaseVerilogParserListener) EnterMos_switchtype(ctx *Mos_switchtypeContext) {}

// ExitMos_switchtype is called when production mos_switchtype is exited.
func (s *BaseVerilogParserListener) ExitMos_switchtype(ctx *Mos_switchtypeContext) {}

// EnterN_input_gatetype is called when production n_input_gatetype is entered.
func (s *BaseVerilogParserListener) EnterN_input_gatetype(ctx *N_input_gatetypeContext) {}

// ExitN_input_gatetype is called when production n_input_gatetype is exited.
func (s *BaseVerilogParserListener) ExitN_input_gatetype(ctx *N_input_gatetypeContext) {}

// EnterN_output_gatetype is called when production n_output_gatetype is entered.
func (s *BaseVerilogParserListener) EnterN_output_gatetype(ctx *N_output_gatetypeContext) {}

// ExitN_output_gatetype is called when production n_output_gatetype is exited.
func (s *BaseVerilogParserListener) ExitN_output_gatetype(ctx *N_output_gatetypeContext) {}

// EnterPass_en_switchtype is called when production pass_en_switchtype is entered.
func (s *BaseVerilogParserListener) EnterPass_en_switchtype(ctx *Pass_en_switchtypeContext) {}

// ExitPass_en_switchtype is called when production pass_en_switchtype is exited.
func (s *BaseVerilogParserListener) ExitPass_en_switchtype(ctx *Pass_en_switchtypeContext) {}

// EnterPass_switchtype is called when production pass_switchtype is entered.
func (s *BaseVerilogParserListener) EnterPass_switchtype(ctx *Pass_switchtypeContext) {}

// ExitPass_switchtype is called when production pass_switchtype is exited.
func (s *BaseVerilogParserListener) ExitPass_switchtype(ctx *Pass_switchtypeContext) {}

// EnterModule_instantiation is called when production module_instantiation is entered.
func (s *BaseVerilogParserListener) EnterModule_instantiation(ctx *Module_instantiationContext) {}

// ExitModule_instantiation is called when production module_instantiation is exited.
func (s *BaseVerilogParserListener) ExitModule_instantiation(ctx *Module_instantiationContext) {}

// EnterParameter_value_assignment is called when production parameter_value_assignment is entered.
func (s *BaseVerilogParserListener) EnterParameter_value_assignment(ctx *Parameter_value_assignmentContext) {
}

// ExitParameter_value_assignment is called when production parameter_value_assignment is exited.
func (s *BaseVerilogParserListener) ExitParameter_value_assignment(ctx *Parameter_value_assignmentContext) {
}

// EnterList_of_parameter_assignments is called when production list_of_parameter_assignments is entered.
func (s *BaseVerilogParserListener) EnterList_of_parameter_assignments(ctx *List_of_parameter_assignmentsContext) {
}

// ExitList_of_parameter_assignments is called when production list_of_parameter_assignments is exited.
func (s *BaseVerilogParserListener) ExitList_of_parameter_assignments(ctx *List_of_parameter_assignmentsContext) {
}

// EnterOrdered_parameter_assignment is called when production ordered_parameter_assignment is entered.
func (s *BaseVerilogParserListener) EnterOrdered_parameter_assignment(ctx *Ordered_parameter_assignmentContext) {
}

// ExitOrdered_parameter_assignment is called when production ordered_parameter_assignment is exited.
func (s *BaseVerilogParserListener) ExitOrdered_parameter_assignment(ctx *Ordered_parameter_assignmentContext) {
}

// EnterNamed_parameter_assignment is called when production named_parameter_assignment is entered.
func (s *BaseVerilogParserListener) EnterNamed_parameter_assignment(ctx *Named_parameter_assignmentContext) {
}

// ExitNamed_parameter_assignment is called when production named_parameter_assignment is exited.
func (s *BaseVerilogParserListener) ExitNamed_parameter_assignment(ctx *Named_parameter_assignmentContext) {
}

// EnterModule_instance is called when production module_instance is entered.
func (s *BaseVerilogParserListener) EnterModule_instance(ctx *Module_instanceContext) {}

// ExitModule_instance is called when production module_instance is exited.
func (s *BaseVerilogParserListener) ExitModule_instance(ctx *Module_instanceContext) {}

// EnterName_of_instance is called when production name_of_instance is entered.
func (s *BaseVerilogParserListener) EnterName_of_instance(ctx *Name_of_instanceContext) {}

// ExitName_of_instance is called when production name_of_instance is exited.
func (s *BaseVerilogParserListener) ExitName_of_instance(ctx *Name_of_instanceContext) {}

// EnterList_of_port_connections is called when production list_of_port_connections is entered.
func (s *BaseVerilogParserListener) EnterList_of_port_connections(ctx *List_of_port_connectionsContext) {
}

// ExitList_of_port_connections is called when production list_of_port_connections is exited.
func (s *BaseVerilogParserListener) ExitList_of_port_connections(ctx *List_of_port_connectionsContext) {
}

// EnterOrdered_port_connection is called when production ordered_port_connection is entered.
func (s *BaseVerilogParserListener) EnterOrdered_port_connection(ctx *Ordered_port_connectionContext) {
}

// ExitOrdered_port_connection is called when production ordered_port_connection is exited.
func (s *BaseVerilogParserListener) ExitOrdered_port_connection(ctx *Ordered_port_connectionContext) {
}

// EnterNamed_port_connection is called when production named_port_connection is entered.
func (s *BaseVerilogParserListener) EnterNamed_port_connection(ctx *Named_port_connectionContext) {}

// ExitNamed_port_connection is called when production named_port_connection is exited.
func (s *BaseVerilogParserListener) ExitNamed_port_connection(ctx *Named_port_connectionContext) {}

// EnterGenerated_instantiation is called when production generated_instantiation is entered.
func (s *BaseVerilogParserListener) EnterGenerated_instantiation(ctx *Generated_instantiationContext) {
}

// ExitGenerated_instantiation is called when production generated_instantiation is exited.
func (s *BaseVerilogParserListener) ExitGenerated_instantiation(ctx *Generated_instantiationContext) {
}

// EnterGenerate_item_or_null is called when production generate_item_or_null is entered.
func (s *BaseVerilogParserListener) EnterGenerate_item_or_null(ctx *Generate_item_or_nullContext) {}

// ExitGenerate_item_or_null is called when production generate_item_or_null is exited.
func (s *BaseVerilogParserListener) ExitGenerate_item_or_null(ctx *Generate_item_or_nullContext) {}

// EnterGenerate_item is called when production generate_item is entered.
func (s *BaseVerilogParserListener) EnterGenerate_item(ctx *Generate_itemContext) {}

// ExitGenerate_item is called when production generate_item is exited.
func (s *BaseVerilogParserListener) ExitGenerate_item(ctx *Generate_itemContext) {}

// EnterGenerate_conditional_statement is called when production generate_conditional_statement is entered.
func (s *BaseVerilogParserListener) EnterGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// ExitGenerate_conditional_statement is called when production generate_conditional_statement is exited.
func (s *BaseVerilogParserListener) ExitGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// EnterGenerate_case_statement is called when production generate_case_statement is entered.
func (s *BaseVerilogParserListener) EnterGenerate_case_statement(ctx *Generate_case_statementContext) {
}

// ExitGenerate_case_statement is called when production generate_case_statement is exited.
func (s *BaseVerilogParserListener) ExitGenerate_case_statement(ctx *Generate_case_statementContext) {
}

// EnterGenvar_case_item is called when production genvar_case_item is entered.
func (s *BaseVerilogParserListener) EnterGenvar_case_item(ctx *Genvar_case_itemContext) {}

// ExitGenvar_case_item is called when production genvar_case_item is exited.
func (s *BaseVerilogParserListener) ExitGenvar_case_item(ctx *Genvar_case_itemContext) {}

// EnterGenerate_loop_statement is called when production generate_loop_statement is entered.
func (s *BaseVerilogParserListener) EnterGenerate_loop_statement(ctx *Generate_loop_statementContext) {
}

// ExitGenerate_loop_statement is called when production generate_loop_statement is exited.
func (s *BaseVerilogParserListener) ExitGenerate_loop_statement(ctx *Generate_loop_statementContext) {
}

// EnterGenvar_assignment is called when production genvar_assignment is entered.
func (s *BaseVerilogParserListener) EnterGenvar_assignment(ctx *Genvar_assignmentContext) {}

// ExitGenvar_assignment is called when production genvar_assignment is exited.
func (s *BaseVerilogParserListener) ExitGenvar_assignment(ctx *Genvar_assignmentContext) {}

// EnterGenerate_block is called when production generate_block is entered.
func (s *BaseVerilogParserListener) EnterGenerate_block(ctx *Generate_blockContext) {}

// ExitGenerate_block is called when production generate_block is exited.
func (s *BaseVerilogParserListener) ExitGenerate_block(ctx *Generate_blockContext) {}

// EnterContinuous_assign is called when production continuous_assign is entered.
func (s *BaseVerilogParserListener) EnterContinuous_assign(ctx *Continuous_assignContext) {}

// ExitContinuous_assign is called when production continuous_assign is exited.
func (s *BaseVerilogParserListener) ExitContinuous_assign(ctx *Continuous_assignContext) {}

// EnterList_of_net_assignments is called when production list_of_net_assignments is entered.
func (s *BaseVerilogParserListener) EnterList_of_net_assignments(ctx *List_of_net_assignmentsContext) {
}

// ExitList_of_net_assignments is called when production list_of_net_assignments is exited.
func (s *BaseVerilogParserListener) ExitList_of_net_assignments(ctx *List_of_net_assignmentsContext) {
}

// EnterNet_assignment is called when production net_assignment is entered.
func (s *BaseVerilogParserListener) EnterNet_assignment(ctx *Net_assignmentContext) {}

// ExitNet_assignment is called when production net_assignment is exited.
func (s *BaseVerilogParserListener) ExitNet_assignment(ctx *Net_assignmentContext) {}

// EnterInitial_construct is called when production initial_construct is entered.
func (s *BaseVerilogParserListener) EnterInitial_construct(ctx *Initial_constructContext) {}

// ExitInitial_construct is called when production initial_construct is exited.
func (s *BaseVerilogParserListener) ExitInitial_construct(ctx *Initial_constructContext) {}

// EnterAlways_construct is called when production always_construct is entered.
func (s *BaseVerilogParserListener) EnterAlways_construct(ctx *Always_constructContext) {}

// ExitAlways_construct is called when production always_construct is exited.
func (s *BaseVerilogParserListener) ExitAlways_construct(ctx *Always_constructContext) {}

// EnterBlocking_assignment is called when production blocking_assignment is entered.
func (s *BaseVerilogParserListener) EnterBlocking_assignment(ctx *Blocking_assignmentContext) {}

// ExitBlocking_assignment is called when production blocking_assignment is exited.
func (s *BaseVerilogParserListener) ExitBlocking_assignment(ctx *Blocking_assignmentContext) {}

// EnterNonblocking_assignment is called when production nonblocking_assignment is entered.
func (s *BaseVerilogParserListener) EnterNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *BaseVerilogParserListener) ExitNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// EnterProcedural_continuous_assignments is called when production procedural_continuous_assignments is entered.
func (s *BaseVerilogParserListener) EnterProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// ExitProcedural_continuous_assignments is called when production procedural_continuous_assignments is exited.
func (s *BaseVerilogParserListener) ExitProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// EnterFunction_blocking_assignment is called when production function_blocking_assignment is entered.
func (s *BaseVerilogParserListener) EnterFunction_blocking_assignment(ctx *Function_blocking_assignmentContext) {
}

// ExitFunction_blocking_assignment is called when production function_blocking_assignment is exited.
func (s *BaseVerilogParserListener) ExitFunction_blocking_assignment(ctx *Function_blocking_assignmentContext) {
}

// EnterFunction_statement_or_null is called when production function_statement_or_null is entered.
func (s *BaseVerilogParserListener) EnterFunction_statement_or_null(ctx *Function_statement_or_nullContext) {
}

// ExitFunction_statement_or_null is called when production function_statement_or_null is exited.
func (s *BaseVerilogParserListener) ExitFunction_statement_or_null(ctx *Function_statement_or_nullContext) {
}

// EnterFunction_seq_block is called when production function_seq_block is entered.
func (s *BaseVerilogParserListener) EnterFunction_seq_block(ctx *Function_seq_blockContext) {}

// ExitFunction_seq_block is called when production function_seq_block is exited.
func (s *BaseVerilogParserListener) ExitFunction_seq_block(ctx *Function_seq_blockContext) {}

// EnterVariable_assignment is called when production variable_assignment is entered.
func (s *BaseVerilogParserListener) EnterVariable_assignment(ctx *Variable_assignmentContext) {}

// ExitVariable_assignment is called when production variable_assignment is exited.
func (s *BaseVerilogParserListener) ExitVariable_assignment(ctx *Variable_assignmentContext) {}

// EnterPar_block is called when production par_block is entered.
func (s *BaseVerilogParserListener) EnterPar_block(ctx *Par_blockContext) {}

// ExitPar_block is called when production par_block is exited.
func (s *BaseVerilogParserListener) ExitPar_block(ctx *Par_blockContext) {}

// EnterSeq_block is called when production seq_block is entered.
func (s *BaseVerilogParserListener) EnterSeq_block(ctx *Seq_blockContext) {}

// ExitSeq_block is called when production seq_block is exited.
func (s *BaseVerilogParserListener) ExitSeq_block(ctx *Seq_blockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseVerilogParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseVerilogParserListener) ExitStatement(ctx *StatementContext) {}

// EnterStatement_or_null is called when production statement_or_null is entered.
func (s *BaseVerilogParserListener) EnterStatement_or_null(ctx *Statement_or_nullContext) {}

// ExitStatement_or_null is called when production statement_or_null is exited.
func (s *BaseVerilogParserListener) ExitStatement_or_null(ctx *Statement_or_nullContext) {}

// EnterFunction_statement is called when production function_statement is entered.
func (s *BaseVerilogParserListener) EnterFunction_statement(ctx *Function_statementContext) {}

// ExitFunction_statement is called when production function_statement is exited.
func (s *BaseVerilogParserListener) ExitFunction_statement(ctx *Function_statementContext) {}

// EnterDelay_or_event_control is called when production delay_or_event_control is entered.
func (s *BaseVerilogParserListener) EnterDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// ExitDelay_or_event_control is called when production delay_or_event_control is exited.
func (s *BaseVerilogParserListener) ExitDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// EnterDelay_control is called when production delay_control is entered.
func (s *BaseVerilogParserListener) EnterDelay_control(ctx *Delay_controlContext) {}

// ExitDelay_control is called when production delay_control is exited.
func (s *BaseVerilogParserListener) ExitDelay_control(ctx *Delay_controlContext) {}

// EnterDisable_statement is called when production disable_statement is entered.
func (s *BaseVerilogParserListener) EnterDisable_statement(ctx *Disable_statementContext) {}

// ExitDisable_statement is called when production disable_statement is exited.
func (s *BaseVerilogParserListener) ExitDisable_statement(ctx *Disable_statementContext) {}

// EnterEvent_control is called when production event_control is entered.
func (s *BaseVerilogParserListener) EnterEvent_control(ctx *Event_controlContext) {}

// ExitEvent_control is called when production event_control is exited.
func (s *BaseVerilogParserListener) ExitEvent_control(ctx *Event_controlContext) {}

// EnterEvent_trigger is called when production event_trigger is entered.
func (s *BaseVerilogParserListener) EnterEvent_trigger(ctx *Event_triggerContext) {}

// ExitEvent_trigger is called when production event_trigger is exited.
func (s *BaseVerilogParserListener) ExitEvent_trigger(ctx *Event_triggerContext) {}

// EnterEvent_expression is called when production event_expression is entered.
func (s *BaseVerilogParserListener) EnterEvent_expression(ctx *Event_expressionContext) {}

// ExitEvent_expression is called when production event_expression is exited.
func (s *BaseVerilogParserListener) ExitEvent_expression(ctx *Event_expressionContext) {}

// EnterEvent_primary is called when production event_primary is entered.
func (s *BaseVerilogParserListener) EnterEvent_primary(ctx *Event_primaryContext) {}

// ExitEvent_primary is called when production event_primary is exited.
func (s *BaseVerilogParserListener) ExitEvent_primary(ctx *Event_primaryContext) {}

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *BaseVerilogParserListener) EnterProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *BaseVerilogParserListener) ExitProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// EnterWait_statement is called when production wait_statement is entered.
func (s *BaseVerilogParserListener) EnterWait_statement(ctx *Wait_statementContext) {}

// ExitWait_statement is called when production wait_statement is exited.
func (s *BaseVerilogParserListener) ExitWait_statement(ctx *Wait_statementContext) {}

// EnterConditional_statement is called when production conditional_statement is entered.
func (s *BaseVerilogParserListener) EnterConditional_statement(ctx *Conditional_statementContext) {}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *BaseVerilogParserListener) ExitConditional_statement(ctx *Conditional_statementContext) {}

// EnterIf_else_if_statement is called when production if_else_if_statement is entered.
func (s *BaseVerilogParserListener) EnterIf_else_if_statement(ctx *If_else_if_statementContext) {}

// ExitIf_else_if_statement is called when production if_else_if_statement is exited.
func (s *BaseVerilogParserListener) ExitIf_else_if_statement(ctx *If_else_if_statementContext) {}

// EnterFunction_conditional_statement is called when production function_conditional_statement is entered.
func (s *BaseVerilogParserListener) EnterFunction_conditional_statement(ctx *Function_conditional_statementContext) {
}

// ExitFunction_conditional_statement is called when production function_conditional_statement is exited.
func (s *BaseVerilogParserListener) ExitFunction_conditional_statement(ctx *Function_conditional_statementContext) {
}

// EnterFunction_if_else_if_statement is called when production function_if_else_if_statement is entered.
func (s *BaseVerilogParserListener) EnterFunction_if_else_if_statement(ctx *Function_if_else_if_statementContext) {
}

// ExitFunction_if_else_if_statement is called when production function_if_else_if_statement is exited.
func (s *BaseVerilogParserListener) ExitFunction_if_else_if_statement(ctx *Function_if_else_if_statementContext) {
}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseVerilogParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseVerilogParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterCase_item is called when production case_item is entered.
func (s *BaseVerilogParserListener) EnterCase_item(ctx *Case_itemContext) {}

// ExitCase_item is called when production case_item is exited.
func (s *BaseVerilogParserListener) ExitCase_item(ctx *Case_itemContext) {}

// EnterFunction_case_statement is called when production function_case_statement is entered.
func (s *BaseVerilogParserListener) EnterFunction_case_statement(ctx *Function_case_statementContext) {
}

// ExitFunction_case_statement is called when production function_case_statement is exited.
func (s *BaseVerilogParserListener) ExitFunction_case_statement(ctx *Function_case_statementContext) {
}

// EnterFunction_case_item is called when production function_case_item is entered.
func (s *BaseVerilogParserListener) EnterFunction_case_item(ctx *Function_case_itemContext) {}

// ExitFunction_case_item is called when production function_case_item is exited.
func (s *BaseVerilogParserListener) ExitFunction_case_item(ctx *Function_case_itemContext) {}

// EnterFunction_loop_statement is called when production function_loop_statement is entered.
func (s *BaseVerilogParserListener) EnterFunction_loop_statement(ctx *Function_loop_statementContext) {
}

// ExitFunction_loop_statement is called when production function_loop_statement is exited.
func (s *BaseVerilogParserListener) ExitFunction_loop_statement(ctx *Function_loop_statementContext) {
}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseVerilogParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseVerilogParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterSystem_task_enable is called when production system_task_enable is entered.
func (s *BaseVerilogParserListener) EnterSystem_task_enable(ctx *System_task_enableContext) {}

// ExitSystem_task_enable is called when production system_task_enable is exited.
func (s *BaseVerilogParserListener) ExitSystem_task_enable(ctx *System_task_enableContext) {}

// EnterTask_enable is called when production task_enable is entered.
func (s *BaseVerilogParserListener) EnterTask_enable(ctx *Task_enableContext) {}

// ExitTask_enable is called when production task_enable is exited.
func (s *BaseVerilogParserListener) ExitTask_enable(ctx *Task_enableContext) {}

// EnterSpecify_block is called when production specify_block is entered.
func (s *BaseVerilogParserListener) EnterSpecify_block(ctx *Specify_blockContext) {}

// ExitSpecify_block is called when production specify_block is exited.
func (s *BaseVerilogParserListener) ExitSpecify_block(ctx *Specify_blockContext) {}

// EnterSpecify_item is called when production specify_item is entered.
func (s *BaseVerilogParserListener) EnterSpecify_item(ctx *Specify_itemContext) {}

// ExitSpecify_item is called when production specify_item is exited.
func (s *BaseVerilogParserListener) ExitSpecify_item(ctx *Specify_itemContext) {}

// EnterPulsestyle_declaration is called when production pulsestyle_declaration is entered.
func (s *BaseVerilogParserListener) EnterPulsestyle_declaration(ctx *Pulsestyle_declarationContext) {}

// ExitPulsestyle_declaration is called when production pulsestyle_declaration is exited.
func (s *BaseVerilogParserListener) ExitPulsestyle_declaration(ctx *Pulsestyle_declarationContext) {}

// EnterShowcancelled_declaration is called when production showcancelled_declaration is entered.
func (s *BaseVerilogParserListener) EnterShowcancelled_declaration(ctx *Showcancelled_declarationContext) {
}

// ExitShowcancelled_declaration is called when production showcancelled_declaration is exited.
func (s *BaseVerilogParserListener) ExitShowcancelled_declaration(ctx *Showcancelled_declarationContext) {
}

// EnterPath_declaration is called when production path_declaration is entered.
func (s *BaseVerilogParserListener) EnterPath_declaration(ctx *Path_declarationContext) {}

// ExitPath_declaration is called when production path_declaration is exited.
func (s *BaseVerilogParserListener) ExitPath_declaration(ctx *Path_declarationContext) {}

// EnterSimple_path_declaration is called when production simple_path_declaration is entered.
func (s *BaseVerilogParserListener) EnterSimple_path_declaration(ctx *Simple_path_declarationContext) {
}

// ExitSimple_path_declaration is called when production simple_path_declaration is exited.
func (s *BaseVerilogParserListener) ExitSimple_path_declaration(ctx *Simple_path_declarationContext) {
}

// EnterParallel_path_description is called when production parallel_path_description is entered.
func (s *BaseVerilogParserListener) EnterParallel_path_description(ctx *Parallel_path_descriptionContext) {
}

// ExitParallel_path_description is called when production parallel_path_description is exited.
func (s *BaseVerilogParserListener) ExitParallel_path_description(ctx *Parallel_path_descriptionContext) {
}

// EnterFull_path_description is called when production full_path_description is entered.
func (s *BaseVerilogParserListener) EnterFull_path_description(ctx *Full_path_descriptionContext) {}

// ExitFull_path_description is called when production full_path_description is exited.
func (s *BaseVerilogParserListener) ExitFull_path_description(ctx *Full_path_descriptionContext) {}

// EnterList_of_path_inputs is called when production list_of_path_inputs is entered.
func (s *BaseVerilogParserListener) EnterList_of_path_inputs(ctx *List_of_path_inputsContext) {}

// ExitList_of_path_inputs is called when production list_of_path_inputs is exited.
func (s *BaseVerilogParserListener) ExitList_of_path_inputs(ctx *List_of_path_inputsContext) {}

// EnterList_of_path_outputs is called when production list_of_path_outputs is entered.
func (s *BaseVerilogParserListener) EnterList_of_path_outputs(ctx *List_of_path_outputsContext) {}

// ExitList_of_path_outputs is called when production list_of_path_outputs is exited.
func (s *BaseVerilogParserListener) ExitList_of_path_outputs(ctx *List_of_path_outputsContext) {}

// EnterSpecify_input_terminal_descriptor is called when production specify_input_terminal_descriptor is entered.
func (s *BaseVerilogParserListener) EnterSpecify_input_terminal_descriptor(ctx *Specify_input_terminal_descriptorContext) {
}

// ExitSpecify_input_terminal_descriptor is called when production specify_input_terminal_descriptor is exited.
func (s *BaseVerilogParserListener) ExitSpecify_input_terminal_descriptor(ctx *Specify_input_terminal_descriptorContext) {
}

// EnterSpecify_output_terminal_descriptor is called when production specify_output_terminal_descriptor is entered.
func (s *BaseVerilogParserListener) EnterSpecify_output_terminal_descriptor(ctx *Specify_output_terminal_descriptorContext) {
}

// ExitSpecify_output_terminal_descriptor is called when production specify_output_terminal_descriptor is exited.
func (s *BaseVerilogParserListener) ExitSpecify_output_terminal_descriptor(ctx *Specify_output_terminal_descriptorContext) {
}

// EnterInput_identifier is called when production input_identifier is entered.
func (s *BaseVerilogParserListener) EnterInput_identifier(ctx *Input_identifierContext) {}

// ExitInput_identifier is called when production input_identifier is exited.
func (s *BaseVerilogParserListener) ExitInput_identifier(ctx *Input_identifierContext) {}

// EnterOutput_identifier is called when production output_identifier is entered.
func (s *BaseVerilogParserListener) EnterOutput_identifier(ctx *Output_identifierContext) {}

// ExitOutput_identifier is called when production output_identifier is exited.
func (s *BaseVerilogParserListener) ExitOutput_identifier(ctx *Output_identifierContext) {}

// EnterPath_delay_value is called when production path_delay_value is entered.
func (s *BaseVerilogParserListener) EnterPath_delay_value(ctx *Path_delay_valueContext) {}

// ExitPath_delay_value is called when production path_delay_value is exited.
func (s *BaseVerilogParserListener) ExitPath_delay_value(ctx *Path_delay_valueContext) {}

// EnterList_of_path_delay_expressions is called when production list_of_path_delay_expressions is entered.
func (s *BaseVerilogParserListener) EnterList_of_path_delay_expressions(ctx *List_of_path_delay_expressionsContext) {
}

// ExitList_of_path_delay_expressions is called when production list_of_path_delay_expressions is exited.
func (s *BaseVerilogParserListener) ExitList_of_path_delay_expressions(ctx *List_of_path_delay_expressionsContext) {
}

// EnterT_path_delay_expression is called when production t_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT_path_delay_expression(ctx *T_path_delay_expressionContext) {
}

// ExitT_path_delay_expression is called when production t_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT_path_delay_expression(ctx *T_path_delay_expressionContext) {
}

// EnterTrise_path_delay_expression is called when production trise_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTrise_path_delay_expression(ctx *Trise_path_delay_expressionContext) {
}

// ExitTrise_path_delay_expression is called when production trise_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTrise_path_delay_expression(ctx *Trise_path_delay_expressionContext) {
}

// EnterTfall_path_delay_expression is called when production tfall_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTfall_path_delay_expression(ctx *Tfall_path_delay_expressionContext) {
}

// ExitTfall_path_delay_expression is called when production tfall_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTfall_path_delay_expression(ctx *Tfall_path_delay_expressionContext) {
}

// EnterTz_path_delay_expression is called when production tz_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTz_path_delay_expression(ctx *Tz_path_delay_expressionContext) {
}

// ExitTz_path_delay_expression is called when production tz_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTz_path_delay_expression(ctx *Tz_path_delay_expressionContext) {
}

// EnterT01_path_delay_expression is called when production t01_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT01_path_delay_expression(ctx *T01_path_delay_expressionContext) {
}

// ExitT01_path_delay_expression is called when production t01_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT01_path_delay_expression(ctx *T01_path_delay_expressionContext) {
}

// EnterT10_path_delay_expression is called when production t10_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT10_path_delay_expression(ctx *T10_path_delay_expressionContext) {
}

// ExitT10_path_delay_expression is called when production t10_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT10_path_delay_expression(ctx *T10_path_delay_expressionContext) {
}

// EnterT0z_path_delay_expression is called when production t0z_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT0z_path_delay_expression(ctx *T0z_path_delay_expressionContext) {
}

// ExitT0z_path_delay_expression is called when production t0z_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT0z_path_delay_expression(ctx *T0z_path_delay_expressionContext) {
}

// EnterTz1_path_delay_expression is called when production tz1_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTz1_path_delay_expression(ctx *Tz1_path_delay_expressionContext) {
}

// ExitTz1_path_delay_expression is called when production tz1_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTz1_path_delay_expression(ctx *Tz1_path_delay_expressionContext) {
}

// EnterT1z_path_delay_expression is called when production t1z_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT1z_path_delay_expression(ctx *T1z_path_delay_expressionContext) {
}

// ExitT1z_path_delay_expression is called when production t1z_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT1z_path_delay_expression(ctx *T1z_path_delay_expressionContext) {
}

// EnterTz0_path_delay_expression is called when production tz0_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTz0_path_delay_expression(ctx *Tz0_path_delay_expressionContext) {
}

// ExitTz0_path_delay_expression is called when production tz0_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTz0_path_delay_expression(ctx *Tz0_path_delay_expressionContext) {
}

// EnterT0x_path_delay_expression is called when production t0x_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT0x_path_delay_expression(ctx *T0x_path_delay_expressionContext) {
}

// ExitT0x_path_delay_expression is called when production t0x_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT0x_path_delay_expression(ctx *T0x_path_delay_expressionContext) {
}

// EnterTx1_path_delay_expression is called when production tx1_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTx1_path_delay_expression(ctx *Tx1_path_delay_expressionContext) {
}

// ExitTx1_path_delay_expression is called when production tx1_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTx1_path_delay_expression(ctx *Tx1_path_delay_expressionContext) {
}

// EnterT1x_path_delay_expression is called when production t1x_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterT1x_path_delay_expression(ctx *T1x_path_delay_expressionContext) {
}

// ExitT1x_path_delay_expression is called when production t1x_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitT1x_path_delay_expression(ctx *T1x_path_delay_expressionContext) {
}

// EnterTx0_path_delay_expression is called when production tx0_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTx0_path_delay_expression(ctx *Tx0_path_delay_expressionContext) {
}

// ExitTx0_path_delay_expression is called when production tx0_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTx0_path_delay_expression(ctx *Tx0_path_delay_expressionContext) {
}

// EnterTxz_path_delay_expression is called when production txz_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTxz_path_delay_expression(ctx *Txz_path_delay_expressionContext) {
}

// ExitTxz_path_delay_expression is called when production txz_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTxz_path_delay_expression(ctx *Txz_path_delay_expressionContext) {
}

// EnterTzx_path_delay_expression is called when production tzx_path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterTzx_path_delay_expression(ctx *Tzx_path_delay_expressionContext) {
}

// ExitTzx_path_delay_expression is called when production tzx_path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitTzx_path_delay_expression(ctx *Tzx_path_delay_expressionContext) {
}

// EnterPath_delay_expression is called when production path_delay_expression is entered.
func (s *BaseVerilogParserListener) EnterPath_delay_expression(ctx *Path_delay_expressionContext) {}

// ExitPath_delay_expression is called when production path_delay_expression is exited.
func (s *BaseVerilogParserListener) ExitPath_delay_expression(ctx *Path_delay_expressionContext) {}

// EnterEdge_sensitive_path_declaration is called when production edge_sensitive_path_declaration is entered.
func (s *BaseVerilogParserListener) EnterEdge_sensitive_path_declaration(ctx *Edge_sensitive_path_declarationContext) {
}

// ExitEdge_sensitive_path_declaration is called when production edge_sensitive_path_declaration is exited.
func (s *BaseVerilogParserListener) ExitEdge_sensitive_path_declaration(ctx *Edge_sensitive_path_declarationContext) {
}

// EnterParallel_edge_sensitive_path_description is called when production parallel_edge_sensitive_path_description is entered.
func (s *BaseVerilogParserListener) EnterParallel_edge_sensitive_path_description(ctx *Parallel_edge_sensitive_path_descriptionContext) {
}

// ExitParallel_edge_sensitive_path_description is called when production parallel_edge_sensitive_path_description is exited.
func (s *BaseVerilogParserListener) ExitParallel_edge_sensitive_path_description(ctx *Parallel_edge_sensitive_path_descriptionContext) {
}

// EnterFull_edge_sensitive_path_description is called when production full_edge_sensitive_path_description is entered.
func (s *BaseVerilogParserListener) EnterFull_edge_sensitive_path_description(ctx *Full_edge_sensitive_path_descriptionContext) {
}

// ExitFull_edge_sensitive_path_description is called when production full_edge_sensitive_path_description is exited.
func (s *BaseVerilogParserListener) ExitFull_edge_sensitive_path_description(ctx *Full_edge_sensitive_path_descriptionContext) {
}

// EnterData_source_expression is called when production data_source_expression is entered.
func (s *BaseVerilogParserListener) EnterData_source_expression(ctx *Data_source_expressionContext) {}

// ExitData_source_expression is called when production data_source_expression is exited.
func (s *BaseVerilogParserListener) ExitData_source_expression(ctx *Data_source_expressionContext) {}

// EnterEdge_identifier is called when production edge_identifier is entered.
func (s *BaseVerilogParserListener) EnterEdge_identifier(ctx *Edge_identifierContext) {}

// ExitEdge_identifier is called when production edge_identifier is exited.
func (s *BaseVerilogParserListener) ExitEdge_identifier(ctx *Edge_identifierContext) {}

// EnterState_dependent_path_declaration is called when production state_dependent_path_declaration is entered.
func (s *BaseVerilogParserListener) EnterState_dependent_path_declaration(ctx *State_dependent_path_declarationContext) {
}

// ExitState_dependent_path_declaration is called when production state_dependent_path_declaration is exited.
func (s *BaseVerilogParserListener) ExitState_dependent_path_declaration(ctx *State_dependent_path_declarationContext) {
}

// EnterPolarity_operator is called when production polarity_operator is entered.
func (s *BaseVerilogParserListener) EnterPolarity_operator(ctx *Polarity_operatorContext) {}

// ExitPolarity_operator is called when production polarity_operator is exited.
func (s *BaseVerilogParserListener) ExitPolarity_operator(ctx *Polarity_operatorContext) {}

// EnterSystem_timing_check is called when production system_timing_check is entered.
func (s *BaseVerilogParserListener) EnterSystem_timing_check(ctx *System_timing_checkContext) {}

// ExitSystem_timing_check is called when production system_timing_check is exited.
func (s *BaseVerilogParserListener) ExitSystem_timing_check(ctx *System_timing_checkContext) {}

// EnterSetup_timing_check is called when production setup_timing_check is entered.
func (s *BaseVerilogParserListener) EnterSetup_timing_check(ctx *Setup_timing_checkContext) {}

// ExitSetup_timing_check is called when production setup_timing_check is exited.
func (s *BaseVerilogParserListener) ExitSetup_timing_check(ctx *Setup_timing_checkContext) {}

// EnterHold_timing_check is called when production hold_timing_check is entered.
func (s *BaseVerilogParserListener) EnterHold_timing_check(ctx *Hold_timing_checkContext) {}

// ExitHold_timing_check is called when production hold_timing_check is exited.
func (s *BaseVerilogParserListener) ExitHold_timing_check(ctx *Hold_timing_checkContext) {}

// EnterSetuphold_timing_check is called when production setuphold_timing_check is entered.
func (s *BaseVerilogParserListener) EnterSetuphold_timing_check(ctx *Setuphold_timing_checkContext) {}

// ExitSetuphold_timing_check is called when production setuphold_timing_check is exited.
func (s *BaseVerilogParserListener) ExitSetuphold_timing_check(ctx *Setuphold_timing_checkContext) {}

// EnterRecovery_timing_check is called when production recovery_timing_check is entered.
func (s *BaseVerilogParserListener) EnterRecovery_timing_check(ctx *Recovery_timing_checkContext) {}

// ExitRecovery_timing_check is called when production recovery_timing_check is exited.
func (s *BaseVerilogParserListener) ExitRecovery_timing_check(ctx *Recovery_timing_checkContext) {}

// EnterRemoval_timing_check is called when production removal_timing_check is entered.
func (s *BaseVerilogParserListener) EnterRemoval_timing_check(ctx *Removal_timing_checkContext) {}

// ExitRemoval_timing_check is called when production removal_timing_check is exited.
func (s *BaseVerilogParserListener) ExitRemoval_timing_check(ctx *Removal_timing_checkContext) {}

// EnterRecrem_timing_check is called when production recrem_timing_check is entered.
func (s *BaseVerilogParserListener) EnterRecrem_timing_check(ctx *Recrem_timing_checkContext) {}

// ExitRecrem_timing_check is called when production recrem_timing_check is exited.
func (s *BaseVerilogParserListener) ExitRecrem_timing_check(ctx *Recrem_timing_checkContext) {}

// EnterSkew_timing_check is called when production skew_timing_check is entered.
func (s *BaseVerilogParserListener) EnterSkew_timing_check(ctx *Skew_timing_checkContext) {}

// ExitSkew_timing_check is called when production skew_timing_check is exited.
func (s *BaseVerilogParserListener) ExitSkew_timing_check(ctx *Skew_timing_checkContext) {}

// EnterTimeskew_timing_check is called when production timeskew_timing_check is entered.
func (s *BaseVerilogParserListener) EnterTimeskew_timing_check(ctx *Timeskew_timing_checkContext) {}

// ExitTimeskew_timing_check is called when production timeskew_timing_check is exited.
func (s *BaseVerilogParserListener) ExitTimeskew_timing_check(ctx *Timeskew_timing_checkContext) {}

// EnterFullskew_timing_check is called when production fullskew_timing_check is entered.
func (s *BaseVerilogParserListener) EnterFullskew_timing_check(ctx *Fullskew_timing_checkContext) {}

// ExitFullskew_timing_check is called when production fullskew_timing_check is exited.
func (s *BaseVerilogParserListener) ExitFullskew_timing_check(ctx *Fullskew_timing_checkContext) {}

// EnterPeriod_timing_check is called when production period_timing_check is entered.
func (s *BaseVerilogParserListener) EnterPeriod_timing_check(ctx *Period_timing_checkContext) {}

// ExitPeriod_timing_check is called when production period_timing_check is exited.
func (s *BaseVerilogParserListener) ExitPeriod_timing_check(ctx *Period_timing_checkContext) {}

// EnterWidth_timing_check is called when production width_timing_check is entered.
func (s *BaseVerilogParserListener) EnterWidth_timing_check(ctx *Width_timing_checkContext) {}

// ExitWidth_timing_check is called when production width_timing_check is exited.
func (s *BaseVerilogParserListener) ExitWidth_timing_check(ctx *Width_timing_checkContext) {}

// EnterNochange_timing_check is called when production nochange_timing_check is entered.
func (s *BaseVerilogParserListener) EnterNochange_timing_check(ctx *Nochange_timing_checkContext) {}

// ExitNochange_timing_check is called when production nochange_timing_check is exited.
func (s *BaseVerilogParserListener) ExitNochange_timing_check(ctx *Nochange_timing_checkContext) {}

// EnterChecktime_condition is called when production checktime_condition is entered.
func (s *BaseVerilogParserListener) EnterChecktime_condition(ctx *Checktime_conditionContext) {}

// ExitChecktime_condition is called when production checktime_condition is exited.
func (s *BaseVerilogParserListener) ExitChecktime_condition(ctx *Checktime_conditionContext) {}

// EnterControlled_reference_event is called when production controlled_reference_event is entered.
func (s *BaseVerilogParserListener) EnterControlled_reference_event(ctx *Controlled_reference_eventContext) {
}

// ExitControlled_reference_event is called when production controlled_reference_event is exited.
func (s *BaseVerilogParserListener) ExitControlled_reference_event(ctx *Controlled_reference_eventContext) {
}

// EnterData_event is called when production data_event is entered.
func (s *BaseVerilogParserListener) EnterData_event(ctx *Data_eventContext) {}

// ExitData_event is called when production data_event is exited.
func (s *BaseVerilogParserListener) ExitData_event(ctx *Data_eventContext) {}

// EnterDelayed_data is called when production delayed_data is entered.
func (s *BaseVerilogParserListener) EnterDelayed_data(ctx *Delayed_dataContext) {}

// ExitDelayed_data is called when production delayed_data is exited.
func (s *BaseVerilogParserListener) ExitDelayed_data(ctx *Delayed_dataContext) {}

// EnterDelayed_reference is called when production delayed_reference is entered.
func (s *BaseVerilogParserListener) EnterDelayed_reference(ctx *Delayed_referenceContext) {}

// ExitDelayed_reference is called when production delayed_reference is exited.
func (s *BaseVerilogParserListener) ExitDelayed_reference(ctx *Delayed_referenceContext) {}

// EnterEnd_edge_offset is called when production end_edge_offset is entered.
func (s *BaseVerilogParserListener) EnterEnd_edge_offset(ctx *End_edge_offsetContext) {}

// ExitEnd_edge_offset is called when production end_edge_offset is exited.
func (s *BaseVerilogParserListener) ExitEnd_edge_offset(ctx *End_edge_offsetContext) {}

// EnterEvent_based_flag is called when production event_based_flag is entered.
func (s *BaseVerilogParserListener) EnterEvent_based_flag(ctx *Event_based_flagContext) {}

// ExitEvent_based_flag is called when production event_based_flag is exited.
func (s *BaseVerilogParserListener) ExitEvent_based_flag(ctx *Event_based_flagContext) {}

// EnterNotify_reg is called when production notify_reg is entered.
func (s *BaseVerilogParserListener) EnterNotify_reg(ctx *Notify_regContext) {}

// ExitNotify_reg is called when production notify_reg is exited.
func (s *BaseVerilogParserListener) ExitNotify_reg(ctx *Notify_regContext) {}

// EnterReference_event is called when production reference_event is entered.
func (s *BaseVerilogParserListener) EnterReference_event(ctx *Reference_eventContext) {}

// ExitReference_event is called when production reference_event is exited.
func (s *BaseVerilogParserListener) ExitReference_event(ctx *Reference_eventContext) {}

// EnterRemain_active_flag is called when production remain_active_flag is entered.
func (s *BaseVerilogParserListener) EnterRemain_active_flag(ctx *Remain_active_flagContext) {}

// ExitRemain_active_flag is called when production remain_active_flag is exited.
func (s *BaseVerilogParserListener) ExitRemain_active_flag(ctx *Remain_active_flagContext) {}

// EnterStamptime_condition is called when production stamptime_condition is entered.
func (s *BaseVerilogParserListener) EnterStamptime_condition(ctx *Stamptime_conditionContext) {}

// ExitStamptime_condition is called when production stamptime_condition is exited.
func (s *BaseVerilogParserListener) ExitStamptime_condition(ctx *Stamptime_conditionContext) {}

// EnterStart_edge_offset is called when production start_edge_offset is entered.
func (s *BaseVerilogParserListener) EnterStart_edge_offset(ctx *Start_edge_offsetContext) {}

// ExitStart_edge_offset is called when production start_edge_offset is exited.
func (s *BaseVerilogParserListener) ExitStart_edge_offset(ctx *Start_edge_offsetContext) {}

// EnterThreshold is called when production threshold is entered.
func (s *BaseVerilogParserListener) EnterThreshold(ctx *ThresholdContext) {}

// ExitThreshold is called when production threshold is exited.
func (s *BaseVerilogParserListener) ExitThreshold(ctx *ThresholdContext) {}

// EnterTiming_check_limit is called when production timing_check_limit is entered.
func (s *BaseVerilogParserListener) EnterTiming_check_limit(ctx *Timing_check_limitContext) {}

// ExitTiming_check_limit is called when production timing_check_limit is exited.
func (s *BaseVerilogParserListener) ExitTiming_check_limit(ctx *Timing_check_limitContext) {}

// EnterTiming_check_event is called when production timing_check_event is entered.
func (s *BaseVerilogParserListener) EnterTiming_check_event(ctx *Timing_check_eventContext) {}

// ExitTiming_check_event is called when production timing_check_event is exited.
func (s *BaseVerilogParserListener) ExitTiming_check_event(ctx *Timing_check_eventContext) {}

// EnterControlled_timing_check_event is called when production controlled_timing_check_event is entered.
func (s *BaseVerilogParserListener) EnterControlled_timing_check_event(ctx *Controlled_timing_check_eventContext) {
}

// ExitControlled_timing_check_event is called when production controlled_timing_check_event is exited.
func (s *BaseVerilogParserListener) ExitControlled_timing_check_event(ctx *Controlled_timing_check_eventContext) {
}

// EnterTiming_check_event_control is called when production timing_check_event_control is entered.
func (s *BaseVerilogParserListener) EnterTiming_check_event_control(ctx *Timing_check_event_controlContext) {
}

// ExitTiming_check_event_control is called when production timing_check_event_control is exited.
func (s *BaseVerilogParserListener) ExitTiming_check_event_control(ctx *Timing_check_event_controlContext) {
}

// EnterSpecify_terminal_descriptor is called when production specify_terminal_descriptor is entered.
func (s *BaseVerilogParserListener) EnterSpecify_terminal_descriptor(ctx *Specify_terminal_descriptorContext) {
}

// ExitSpecify_terminal_descriptor is called when production specify_terminal_descriptor is exited.
func (s *BaseVerilogParserListener) ExitSpecify_terminal_descriptor(ctx *Specify_terminal_descriptorContext) {
}

// EnterEdge_control_specifier is called when production edge_control_specifier is entered.
func (s *BaseVerilogParserListener) EnterEdge_control_specifier(ctx *Edge_control_specifierContext) {}

// ExitEdge_control_specifier is called when production edge_control_specifier is exited.
func (s *BaseVerilogParserListener) ExitEdge_control_specifier(ctx *Edge_control_specifierContext) {}

// EnterTiming_check_condition is called when production timing_check_condition is entered.
func (s *BaseVerilogParserListener) EnterTiming_check_condition(ctx *Timing_check_conditionContext) {}

// ExitTiming_check_condition is called when production timing_check_condition is exited.
func (s *BaseVerilogParserListener) ExitTiming_check_condition(ctx *Timing_check_conditionContext) {}

// EnterScalar_timing_check_condition is called when production scalar_timing_check_condition is entered.
func (s *BaseVerilogParserListener) EnterScalar_timing_check_condition(ctx *Scalar_timing_check_conditionContext) {
}

// ExitScalar_timing_check_condition is called when production scalar_timing_check_condition is exited.
func (s *BaseVerilogParserListener) ExitScalar_timing_check_condition(ctx *Scalar_timing_check_conditionContext) {
}

// EnterScalar_constant is called when production scalar_constant is entered.
func (s *BaseVerilogParserListener) EnterScalar_constant(ctx *Scalar_constantContext) {}

// ExitScalar_constant is called when production scalar_constant is exited.
func (s *BaseVerilogParserListener) ExitScalar_constant(ctx *Scalar_constantContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseVerilogParserListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseVerilogParserListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterConstant_concatenation is called when production constant_concatenation is entered.
func (s *BaseVerilogParserListener) EnterConstant_concatenation(ctx *Constant_concatenationContext) {}

// ExitConstant_concatenation is called when production constant_concatenation is exited.
func (s *BaseVerilogParserListener) ExitConstant_concatenation(ctx *Constant_concatenationContext) {}

// EnterConstant_multiple_concatenation is called when production constant_multiple_concatenation is entered.
func (s *BaseVerilogParserListener) EnterConstant_multiple_concatenation(ctx *Constant_multiple_concatenationContext) {
}

// ExitConstant_multiple_concatenation is called when production constant_multiple_concatenation is exited.
func (s *BaseVerilogParserListener) ExitConstant_multiple_concatenation(ctx *Constant_multiple_concatenationContext) {
}

// EnterModule_path_concatenation is called when production module_path_concatenation is entered.
func (s *BaseVerilogParserListener) EnterModule_path_concatenation(ctx *Module_path_concatenationContext) {
}

// ExitModule_path_concatenation is called when production module_path_concatenation is exited.
func (s *BaseVerilogParserListener) ExitModule_path_concatenation(ctx *Module_path_concatenationContext) {
}

// EnterModule_path_multiple_concatenation is called when production module_path_multiple_concatenation is entered.
func (s *BaseVerilogParserListener) EnterModule_path_multiple_concatenation(ctx *Module_path_multiple_concatenationContext) {
}

// ExitModule_path_multiple_concatenation is called when production module_path_multiple_concatenation is exited.
func (s *BaseVerilogParserListener) ExitModule_path_multiple_concatenation(ctx *Module_path_multiple_concatenationContext) {
}

// EnterMultiple_concatenation is called when production multiple_concatenation is entered.
func (s *BaseVerilogParserListener) EnterMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// ExitMultiple_concatenation is called when production multiple_concatenation is exited.
func (s *BaseVerilogParserListener) ExitMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// EnterNet_concatenation is called when production net_concatenation is entered.
func (s *BaseVerilogParserListener) EnterNet_concatenation(ctx *Net_concatenationContext) {}

// ExitNet_concatenation is called when production net_concatenation is exited.
func (s *BaseVerilogParserListener) ExitNet_concatenation(ctx *Net_concatenationContext) {}

// EnterNet_concatenation_value is called when production net_concatenation_value is entered.
func (s *BaseVerilogParserListener) EnterNet_concatenation_value(ctx *Net_concatenation_valueContext) {
}

// ExitNet_concatenation_value is called when production net_concatenation_value is exited.
func (s *BaseVerilogParserListener) ExitNet_concatenation_value(ctx *Net_concatenation_valueContext) {
}

// EnterVariable_concatenation is called when production variable_concatenation is entered.
func (s *BaseVerilogParserListener) EnterVariable_concatenation(ctx *Variable_concatenationContext) {}

// ExitVariable_concatenation is called when production variable_concatenation is exited.
func (s *BaseVerilogParserListener) ExitVariable_concatenation(ctx *Variable_concatenationContext) {}

// EnterVariable_concatenation_value is called when production variable_concatenation_value is entered.
func (s *BaseVerilogParserListener) EnterVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// ExitVariable_concatenation_value is called when production variable_concatenation_value is exited.
func (s *BaseVerilogParserListener) ExitVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// EnterConstant_function_call is called when production constant_function_call is entered.
func (s *BaseVerilogParserListener) EnterConstant_function_call(ctx *Constant_function_callContext) {}

// ExitConstant_function_call is called when production constant_function_call is exited.
func (s *BaseVerilogParserListener) ExitConstant_function_call(ctx *Constant_function_callContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseVerilogParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseVerilogParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterSystem_function_call is called when production system_function_call is entered.
func (s *BaseVerilogParserListener) EnterSystem_function_call(ctx *System_function_callContext) {}

// ExitSystem_function_call is called when production system_function_call is exited.
func (s *BaseVerilogParserListener) ExitSystem_function_call(ctx *System_function_callContext) {}

// EnterGenvar_function_call is called when production genvar_function_call is entered.
func (s *BaseVerilogParserListener) EnterGenvar_function_call(ctx *Genvar_function_callContext) {}

// ExitGenvar_function_call is called when production genvar_function_call is exited.
func (s *BaseVerilogParserListener) ExitGenvar_function_call(ctx *Genvar_function_callContext) {}

// EnterBase_expression is called when production base_expression is entered.
func (s *BaseVerilogParserListener) EnterBase_expression(ctx *Base_expressionContext) {}

// ExitBase_expression is called when production base_expression is exited.
func (s *BaseVerilogParserListener) ExitBase_expression(ctx *Base_expressionContext) {}

// EnterConstant_base_expression is called when production constant_base_expression is entered.
func (s *BaseVerilogParserListener) EnterConstant_base_expression(ctx *Constant_base_expressionContext) {
}

// ExitConstant_base_expression is called when production constant_base_expression is exited.
func (s *BaseVerilogParserListener) ExitConstant_base_expression(ctx *Constant_base_expressionContext) {
}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseVerilogParserListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseVerilogParserListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterConstant_mintypmax_expression is called when production constant_mintypmax_expression is entered.
func (s *BaseVerilogParserListener) EnterConstant_mintypmax_expression(ctx *Constant_mintypmax_expressionContext) {
}

// ExitConstant_mintypmax_expression is called when production constant_mintypmax_expression is exited.
func (s *BaseVerilogParserListener) ExitConstant_mintypmax_expression(ctx *Constant_mintypmax_expressionContext) {
}

// EnterConstant_range_expression is called when production constant_range_expression is entered.
func (s *BaseVerilogParserListener) EnterConstant_range_expression(ctx *Constant_range_expressionContext) {
}

// ExitConstant_range_expression is called when production constant_range_expression is exited.
func (s *BaseVerilogParserListener) ExitConstant_range_expression(ctx *Constant_range_expressionContext) {
}

// EnterDimension_constant_expression is called when production dimension_constant_expression is entered.
func (s *BaseVerilogParserListener) EnterDimension_constant_expression(ctx *Dimension_constant_expressionContext) {
}

// ExitDimension_constant_expression is called when production dimension_constant_expression is exited.
func (s *BaseVerilogParserListener) ExitDimension_constant_expression(ctx *Dimension_constant_expressionContext) {
}

// EnterExpression is called when production expression is entered.
func (s *BaseVerilogParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVerilogParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseVerilogParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseVerilogParserListener) ExitTerm(ctx *TermContext) {}

// EnterLsb_constant_expression is called when production lsb_constant_expression is entered.
func (s *BaseVerilogParserListener) EnterLsb_constant_expression(ctx *Lsb_constant_expressionContext) {
}

// ExitLsb_constant_expression is called when production lsb_constant_expression is exited.
func (s *BaseVerilogParserListener) ExitLsb_constant_expression(ctx *Lsb_constant_expressionContext) {
}

// EnterMintypmax_expression is called when production mintypmax_expression is entered.
func (s *BaseVerilogParserListener) EnterMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// ExitMintypmax_expression is called when production mintypmax_expression is exited.
func (s *BaseVerilogParserListener) ExitMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// EnterModule_path_conditional_expression is called when production module_path_conditional_expression is entered.
func (s *BaseVerilogParserListener) EnterModule_path_conditional_expression(ctx *Module_path_conditional_expressionContext) {
}

// ExitModule_path_conditional_expression is called when production module_path_conditional_expression is exited.
func (s *BaseVerilogParserListener) ExitModule_path_conditional_expression(ctx *Module_path_conditional_expressionContext) {
}

// EnterModule_path_expression is called when production module_path_expression is entered.
func (s *BaseVerilogParserListener) EnterModule_path_expression(ctx *Module_path_expressionContext) {}

// ExitModule_path_expression is called when production module_path_expression is exited.
func (s *BaseVerilogParserListener) ExitModule_path_expression(ctx *Module_path_expressionContext) {}

// EnterModule_path_mintypmax_expression is called when production module_path_mintypmax_expression is entered.
func (s *BaseVerilogParserListener) EnterModule_path_mintypmax_expression(ctx *Module_path_mintypmax_expressionContext) {
}

// ExitModule_path_mintypmax_expression is called when production module_path_mintypmax_expression is exited.
func (s *BaseVerilogParserListener) ExitModule_path_mintypmax_expression(ctx *Module_path_mintypmax_expressionContext) {
}

// EnterMsb_constant_expression is called when production msb_constant_expression is entered.
func (s *BaseVerilogParserListener) EnterMsb_constant_expression(ctx *Msb_constant_expressionContext) {
}

// ExitMsb_constant_expression is called when production msb_constant_expression is exited.
func (s *BaseVerilogParserListener) ExitMsb_constant_expression(ctx *Msb_constant_expressionContext) {
}

// EnterRange_expression is called when production range_expression is entered.
func (s *BaseVerilogParserListener) EnterRange_expression(ctx *Range_expressionContext) {}

// ExitRange_expression is called when production range_expression is exited.
func (s *BaseVerilogParserListener) ExitRange_expression(ctx *Range_expressionContext) {}

// EnterWidth_constant_expression is called when production width_constant_expression is entered.
func (s *BaseVerilogParserListener) EnterWidth_constant_expression(ctx *Width_constant_expressionContext) {
}

// ExitWidth_constant_expression is called when production width_constant_expression is exited.
func (s *BaseVerilogParserListener) ExitWidth_constant_expression(ctx *Width_constant_expressionContext) {
}

// EnterConstant_primary is called when production constant_primary is entered.
func (s *BaseVerilogParserListener) EnterConstant_primary(ctx *Constant_primaryContext) {}

// ExitConstant_primary is called when production constant_primary is exited.
func (s *BaseVerilogParserListener) ExitConstant_primary(ctx *Constant_primaryContext) {}

// EnterModule_path_primary is called when production module_path_primary is entered.
func (s *BaseVerilogParserListener) EnterModule_path_primary(ctx *Module_path_primaryContext) {}

// ExitModule_path_primary is called when production module_path_primary is exited.
func (s *BaseVerilogParserListener) ExitModule_path_primary(ctx *Module_path_primaryContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseVerilogParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseVerilogParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterNet_lvalue is called when production net_lvalue is entered.
func (s *BaseVerilogParserListener) EnterNet_lvalue(ctx *Net_lvalueContext) {}

// ExitNet_lvalue is called when production net_lvalue is exited.
func (s *BaseVerilogParserListener) ExitNet_lvalue(ctx *Net_lvalueContext) {}

// EnterVariable_lvalue is called when production variable_lvalue is entered.
func (s *BaseVerilogParserListener) EnterVariable_lvalue(ctx *Variable_lvalueContext) {}

// ExitVariable_lvalue is called when production variable_lvalue is exited.
func (s *BaseVerilogParserListener) ExitVariable_lvalue(ctx *Variable_lvalueContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseVerilogParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseVerilogParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseVerilogParserListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseVerilogParserListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_module_path_operator is called when production unary_module_path_operator is entered.
func (s *BaseVerilogParserListener) EnterUnary_module_path_operator(ctx *Unary_module_path_operatorContext) {
}

// ExitUnary_module_path_operator is called when production unary_module_path_operator is exited.
func (s *BaseVerilogParserListener) ExitUnary_module_path_operator(ctx *Unary_module_path_operatorContext) {
}

// EnterBinary_module_path_operator is called when production binary_module_path_operator is entered.
func (s *BaseVerilogParserListener) EnterBinary_module_path_operator(ctx *Binary_module_path_operatorContext) {
}

// ExitBinary_module_path_operator is called when production binary_module_path_operator is exited.
func (s *BaseVerilogParserListener) ExitBinary_module_path_operator(ctx *Binary_module_path_operatorContext) {
}

// EnterNumber is called when production number is entered.
func (s *BaseVerilogParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseVerilogParserListener) ExitNumber(ctx *NumberContext) {}

// EnterAttribute_instance is called when production attribute_instance is entered.
func (s *BaseVerilogParserListener) EnterAttribute_instance(ctx *Attribute_instanceContext) {}

// ExitAttribute_instance is called when production attribute_instance is exited.
func (s *BaseVerilogParserListener) ExitAttribute_instance(ctx *Attribute_instanceContext) {}

// EnterAttr_spec is called when production attr_spec is entered.
func (s *BaseVerilogParserListener) EnterAttr_spec(ctx *Attr_specContext) {}

// ExitAttr_spec is called when production attr_spec is exited.
func (s *BaseVerilogParserListener) ExitAttr_spec(ctx *Attr_specContext) {}

// EnterAttr_name is called when production attr_name is entered.
func (s *BaseVerilogParserListener) EnterAttr_name(ctx *Attr_nameContext) {}

// ExitAttr_name is called when production attr_name is exited.
func (s *BaseVerilogParserListener) ExitAttr_name(ctx *Attr_nameContext) {}

// EnterArrayed_identifier is called when production arrayed_identifier is entered.
func (s *BaseVerilogParserListener) EnterArrayed_identifier(ctx *Arrayed_identifierContext) {}

// ExitArrayed_identifier is called when production arrayed_identifier is exited.
func (s *BaseVerilogParserListener) ExitArrayed_identifier(ctx *Arrayed_identifierContext) {}

// EnterBlock_identifier is called when production block_identifier is entered.
func (s *BaseVerilogParserListener) EnterBlock_identifier(ctx *Block_identifierContext) {}

// ExitBlock_identifier is called when production block_identifier is exited.
func (s *BaseVerilogParserListener) ExitBlock_identifier(ctx *Block_identifierContext) {}

// EnterCell_identifier is called when production cell_identifier is entered.
func (s *BaseVerilogParserListener) EnterCell_identifier(ctx *Cell_identifierContext) {}

// ExitCell_identifier is called when production cell_identifier is exited.
func (s *BaseVerilogParserListener) ExitCell_identifier(ctx *Cell_identifierContext) {}

// EnterConfig_identifier is called when production config_identifier is entered.
func (s *BaseVerilogParserListener) EnterConfig_identifier(ctx *Config_identifierContext) {}

// ExitConfig_identifier is called when production config_identifier is exited.
func (s *BaseVerilogParserListener) ExitConfig_identifier(ctx *Config_identifierContext) {}

// EnterEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is entered.
func (s *BaseVerilogParserListener) EnterEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// ExitEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is exited.
func (s *BaseVerilogParserListener) ExitEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// EnterEscaped_hierarchical_identifier is called when production escaped_hierarchical_identifier is entered.
func (s *BaseVerilogParserListener) EnterEscaped_hierarchical_identifier(ctx *Escaped_hierarchical_identifierContext) {
}

// ExitEscaped_hierarchical_identifier is called when production escaped_hierarchical_identifier is exited.
func (s *BaseVerilogParserListener) ExitEscaped_hierarchical_identifier(ctx *Escaped_hierarchical_identifierContext) {
}

// EnterEvent_identifier is called when production event_identifier is entered.
func (s *BaseVerilogParserListener) EnterEvent_identifier(ctx *Event_identifierContext) {}

// ExitEvent_identifier is called when production event_identifier is exited.
func (s *BaseVerilogParserListener) ExitEvent_identifier(ctx *Event_identifierContext) {}

// EnterFunction_identifier is called when production function_identifier is entered.
func (s *BaseVerilogParserListener) EnterFunction_identifier(ctx *Function_identifierContext) {}

// ExitFunction_identifier is called when production function_identifier is exited.
func (s *BaseVerilogParserListener) ExitFunction_identifier(ctx *Function_identifierContext) {}

// EnterGate_instance_identifier is called when production gate_instance_identifier is entered.
func (s *BaseVerilogParserListener) EnterGate_instance_identifier(ctx *Gate_instance_identifierContext) {
}

// ExitGate_instance_identifier is called when production gate_instance_identifier is exited.
func (s *BaseVerilogParserListener) ExitGate_instance_identifier(ctx *Gate_instance_identifierContext) {
}

// EnterGenerate_block_identifier is called when production generate_block_identifier is entered.
func (s *BaseVerilogParserListener) EnterGenerate_block_identifier(ctx *Generate_block_identifierContext) {
}

// ExitGenerate_block_identifier is called when production generate_block_identifier is exited.
func (s *BaseVerilogParserListener) ExitGenerate_block_identifier(ctx *Generate_block_identifierContext) {
}

// EnterGenvar_function_identifier is called when production genvar_function_identifier is entered.
func (s *BaseVerilogParserListener) EnterGenvar_function_identifier(ctx *Genvar_function_identifierContext) {
}

// ExitGenvar_function_identifier is called when production genvar_function_identifier is exited.
func (s *BaseVerilogParserListener) ExitGenvar_function_identifier(ctx *Genvar_function_identifierContext) {
}

// EnterGenvar_identifier is called when production genvar_identifier is entered.
func (s *BaseVerilogParserListener) EnterGenvar_identifier(ctx *Genvar_identifierContext) {}

// ExitGenvar_identifier is called when production genvar_identifier is exited.
func (s *BaseVerilogParserListener) ExitGenvar_identifier(ctx *Genvar_identifierContext) {}

// EnterHierarchical_block_identifier is called when production hierarchical_block_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// ExitHierarchical_block_identifier is called when production hierarchical_block_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// EnterHierarchical_event_identifier is called when production hierarchical_event_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// ExitHierarchical_event_identifier is called when production hierarchical_event_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// EnterHierarchical_function_identifier is called when production hierarchical_function_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// ExitHierarchical_function_identifier is called when production hierarchical_function_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// EnterHierarchical_identifier is called when production hierarchical_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_identifier(ctx *Hierarchical_identifierContext) {
}

// ExitHierarchical_identifier is called when production hierarchical_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_identifier(ctx *Hierarchical_identifierContext) {
}

// EnterHierarchical_net_identifier is called when production hierarchical_net_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_net_identifier(ctx *Hierarchical_net_identifierContext) {
}

// ExitHierarchical_net_identifier is called when production hierarchical_net_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_net_identifier(ctx *Hierarchical_net_identifierContext) {
}

// EnterHierarchical_variable_identifier is called when production hierarchical_variable_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// ExitHierarchical_variable_identifier is called when production hierarchical_variable_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// EnterHierarchical_task_identifier is called when production hierarchical_task_identifier is entered.
func (s *BaseVerilogParserListener) EnterHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// ExitHierarchical_task_identifier is called when production hierarchical_task_identifier is exited.
func (s *BaseVerilogParserListener) ExitHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseVerilogParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseVerilogParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterInout_port_identifier is called when production inout_port_identifier is entered.
func (s *BaseVerilogParserListener) EnterInout_port_identifier(ctx *Inout_port_identifierContext) {}

// ExitInout_port_identifier is called when production inout_port_identifier is exited.
func (s *BaseVerilogParserListener) ExitInout_port_identifier(ctx *Inout_port_identifierContext) {}

// EnterInput_port_identifier is called when production input_port_identifier is entered.
func (s *BaseVerilogParserListener) EnterInput_port_identifier(ctx *Input_port_identifierContext) {}

// ExitInput_port_identifier is called when production input_port_identifier is exited.
func (s *BaseVerilogParserListener) ExitInput_port_identifier(ctx *Input_port_identifierContext) {}

// EnterInstance_identifier is called when production instance_identifier is entered.
func (s *BaseVerilogParserListener) EnterInstance_identifier(ctx *Instance_identifierContext) {}

// ExitInstance_identifier is called when production instance_identifier is exited.
func (s *BaseVerilogParserListener) ExitInstance_identifier(ctx *Instance_identifierContext) {}

// EnterLibrary_identifier is called when production library_identifier is entered.
func (s *BaseVerilogParserListener) EnterLibrary_identifier(ctx *Library_identifierContext) {}

// ExitLibrary_identifier is called when production library_identifier is exited.
func (s *BaseVerilogParserListener) ExitLibrary_identifier(ctx *Library_identifierContext) {}

// EnterMemory_identifier is called when production memory_identifier is entered.
func (s *BaseVerilogParserListener) EnterMemory_identifier(ctx *Memory_identifierContext) {}

// ExitMemory_identifier is called when production memory_identifier is exited.
func (s *BaseVerilogParserListener) ExitMemory_identifier(ctx *Memory_identifierContext) {}

// EnterModule_identifier is called when production module_identifier is entered.
func (s *BaseVerilogParserListener) EnterModule_identifier(ctx *Module_identifierContext) {}

// ExitModule_identifier is called when production module_identifier is exited.
func (s *BaseVerilogParserListener) ExitModule_identifier(ctx *Module_identifierContext) {}

// EnterModule_instance_identifier is called when production module_instance_identifier is entered.
func (s *BaseVerilogParserListener) EnterModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// ExitModule_instance_identifier is called when production module_instance_identifier is exited.
func (s *BaseVerilogParserListener) ExitModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// EnterNet_identifier is called when production net_identifier is entered.
func (s *BaseVerilogParserListener) EnterNet_identifier(ctx *Net_identifierContext) {}

// ExitNet_identifier is called when production net_identifier is exited.
func (s *BaseVerilogParserListener) ExitNet_identifier(ctx *Net_identifierContext) {}

// EnterOutput_port_identifier is called when production output_port_identifier is entered.
func (s *BaseVerilogParserListener) EnterOutput_port_identifier(ctx *Output_port_identifierContext) {}

// ExitOutput_port_identifier is called when production output_port_identifier is exited.
func (s *BaseVerilogParserListener) ExitOutput_port_identifier(ctx *Output_port_identifierContext) {}

// EnterParameter_identifier is called when production parameter_identifier is entered.
func (s *BaseVerilogParserListener) EnterParameter_identifier(ctx *Parameter_identifierContext) {}

// ExitParameter_identifier is called when production parameter_identifier is exited.
func (s *BaseVerilogParserListener) ExitParameter_identifier(ctx *Parameter_identifierContext) {}

// EnterPort_identifier is called when production port_identifier is entered.
func (s *BaseVerilogParserListener) EnterPort_identifier(ctx *Port_identifierContext) {}

// ExitPort_identifier is called when production port_identifier is exited.
func (s *BaseVerilogParserListener) ExitPort_identifier(ctx *Port_identifierContext) {}

// EnterReal_identifier is called when production real_identifier is entered.
func (s *BaseVerilogParserListener) EnterReal_identifier(ctx *Real_identifierContext) {}

// ExitReal_identifier is called when production real_identifier is exited.
func (s *BaseVerilogParserListener) ExitReal_identifier(ctx *Real_identifierContext) {}

// EnterSimple_arrayed_identifier is called when production simple_arrayed_identifier is entered.
func (s *BaseVerilogParserListener) EnterSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {
}

// ExitSimple_arrayed_identifier is called when production simple_arrayed_identifier is exited.
func (s *BaseVerilogParserListener) ExitSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {
}

// EnterSimple_hierarchical_identifier is called when production simple_hierarchical_identifier is entered.
func (s *BaseVerilogParserListener) EnterSimple_hierarchical_identifier(ctx *Simple_hierarchical_identifierContext) {
}

// ExitSimple_hierarchical_identifier is called when production simple_hierarchical_identifier is exited.
func (s *BaseVerilogParserListener) ExitSimple_hierarchical_identifier(ctx *Simple_hierarchical_identifierContext) {
}

// EnterSpecparam_identifier is called when production specparam_identifier is entered.
func (s *BaseVerilogParserListener) EnterSpecparam_identifier(ctx *Specparam_identifierContext) {}

// ExitSpecparam_identifier is called when production specparam_identifier is exited.
func (s *BaseVerilogParserListener) ExitSpecparam_identifier(ctx *Specparam_identifierContext) {}

// EnterSystem_function_identifier is called when production system_function_identifier is entered.
func (s *BaseVerilogParserListener) EnterSystem_function_identifier(ctx *System_function_identifierContext) {
}

// ExitSystem_function_identifier is called when production system_function_identifier is exited.
func (s *BaseVerilogParserListener) ExitSystem_function_identifier(ctx *System_function_identifierContext) {
}

// EnterSystem_task_identifier is called when production system_task_identifier is entered.
func (s *BaseVerilogParserListener) EnterSystem_task_identifier(ctx *System_task_identifierContext) {}

// ExitSystem_task_identifier is called when production system_task_identifier is exited.
func (s *BaseVerilogParserListener) ExitSystem_task_identifier(ctx *System_task_identifierContext) {}

// EnterTask_identifier is called when production task_identifier is entered.
func (s *BaseVerilogParserListener) EnterTask_identifier(ctx *Task_identifierContext) {}

// ExitTask_identifier is called when production task_identifier is exited.
func (s *BaseVerilogParserListener) ExitTask_identifier(ctx *Task_identifierContext) {}

// EnterTerminal_identifier is called when production terminal_identifier is entered.
func (s *BaseVerilogParserListener) EnterTerminal_identifier(ctx *Terminal_identifierContext) {}

// ExitTerminal_identifier is called when production terminal_identifier is exited.
func (s *BaseVerilogParserListener) ExitTerminal_identifier(ctx *Terminal_identifierContext) {}

// EnterText_macro_identifier is called when production text_macro_identifier is entered.
func (s *BaseVerilogParserListener) EnterText_macro_identifier(ctx *Text_macro_identifierContext) {}

// ExitText_macro_identifier is called when production text_macro_identifier is exited.
func (s *BaseVerilogParserListener) ExitText_macro_identifier(ctx *Text_macro_identifierContext) {}

// EnterTopmodule_identifier is called when production topmodule_identifier is entered.
func (s *BaseVerilogParserListener) EnterTopmodule_identifier(ctx *Topmodule_identifierContext) {}

// ExitTopmodule_identifier is called when production topmodule_identifier is exited.
func (s *BaseVerilogParserListener) ExitTopmodule_identifier(ctx *Topmodule_identifierContext) {}

// EnterUdp_identifier is called when production udp_identifier is entered.
func (s *BaseVerilogParserListener) EnterUdp_identifier(ctx *Udp_identifierContext) {}

// ExitUdp_identifier is called when production udp_identifier is exited.
func (s *BaseVerilogParserListener) ExitUdp_identifier(ctx *Udp_identifierContext) {}

// EnterUdp_instance_identifier is called when production udp_instance_identifier is entered.
func (s *BaseVerilogParserListener) EnterUdp_instance_identifier(ctx *Udp_instance_identifierContext) {
}

// ExitUdp_instance_identifier is called when production udp_instance_identifier is exited.
func (s *BaseVerilogParserListener) ExitUdp_instance_identifier(ctx *Udp_instance_identifierContext) {
}

// EnterVariable_identifier is called when production variable_identifier is entered.
func (s *BaseVerilogParserListener) EnterVariable_identifier(ctx *Variable_identifierContext) {}

// ExitVariable_identifier is called when production variable_identifier is exited.
func (s *BaseVerilogParserListener) ExitVariable_identifier(ctx *Variable_identifierContext) {}

// EnterSimple_hierarchical_branch is called when production simple_hierarchical_branch is entered.
func (s *BaseVerilogParserListener) EnterSimple_hierarchical_branch(ctx *Simple_hierarchical_branchContext) {
}

// ExitSimple_hierarchical_branch is called when production simple_hierarchical_branch is exited.
func (s *BaseVerilogParserListener) ExitSimple_hierarchical_branch(ctx *Simple_hierarchical_branchContext) {
}

// EnterEscaped_hierarchical_branch is called when production escaped_hierarchical_branch is entered.
func (s *BaseVerilogParserListener) EnterEscaped_hierarchical_branch(ctx *Escaped_hierarchical_branchContext) {
}

// ExitEscaped_hierarchical_branch is called when production escaped_hierarchical_branch is exited.
func (s *BaseVerilogParserListener) ExitEscaped_hierarchical_branch(ctx *Escaped_hierarchical_branchContext) {
}

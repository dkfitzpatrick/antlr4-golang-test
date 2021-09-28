// Code generated from SV2017Parser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package sv2017 // SV2017Parser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SV2017ParserListener is a complete listener for a parse tree produced by SV2017Parser.
type SV2017ParserListener interface {
	antlr.ParseTreeListener

	// EnterSource_text is called when entering the source_text production.
	EnterSource_text(c *Source_textContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterHeader_text is called when entering the header_text production.
	EnterHeader_text(c *Header_textContext)

	// EnterDesign_attribute is called when entering the design_attribute production.
	EnterDesign_attribute(c *Design_attributeContext)

	// EnterCompiler_directive is called when entering the compiler_directive production.
	EnterCompiler_directive(c *Compiler_directiveContext)

	// EnterTime_num is called when entering the time_num production.
	EnterTime_num(c *Time_numContext)

	// EnterTime_lit is called when entering the time_lit production.
	EnterTime_lit(c *Time_litContext)

	// EnterTimescale_compiler_directive is called when entering the timescale_compiler_directive production.
	EnterTimescale_compiler_directive(c *Timescale_compiler_directiveContext)

	// EnterDefault_nettype_statement is called when entering the default_nettype_statement production.
	EnterDefault_nettype_statement(c *Default_nettype_statementContext)

	// EnterInclude_svh is called when entering the include_svh production.
	EnterInclude_svh(c *Include_svhContext)

	// EnterAssignment_operator is called when entering the assignment_operator production.
	EnterAssignment_operator(c *Assignment_operatorContext)

	// EnterEdge_identifier is called when entering the edge_identifier production.
	EnterEdge_identifier(c *Edge_identifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterInteger_type is called when entering the integer_type production.
	EnterInteger_type(c *Integer_typeContext)

	// EnterInteger_atom_type is called when entering the integer_atom_type production.
	EnterInteger_atom_type(c *Integer_atom_typeContext)

	// EnterInteger_vector_type is called when entering the integer_vector_type production.
	EnterInteger_vector_type(c *Integer_vector_typeContext)

	// EnterNon_integer_type is called when entering the non_integer_type production.
	EnterNon_integer_type(c *Non_integer_typeContext)

	// EnterNet_type is called when entering the net_type production.
	EnterNet_type(c *Net_typeContext)

	// EnterUnary_module_path_operator is called when entering the unary_module_path_operator production.
	EnterUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterInc_or_dec_operator is called when entering the inc_or_dec_operator production.
	EnterInc_or_dec_operator(c *Inc_or_dec_operatorContext)

	// EnterImplicit_class_handle is called when entering the implicit_class_handle production.
	EnterImplicit_class_handle(c *Implicit_class_handleContext)

	// EnterIntegral_number is called when entering the integral_number production.
	EnterIntegral_number(c *Integral_numberContext)

	// EnterReal_number is called when entering the real_number production.
	EnterReal_number(c *Real_numberContext)

	// EnterAny_system_tf_identifier is called when entering the any_system_tf_identifier production.
	EnterAny_system_tf_identifier(c *Any_system_tf_identifierContext)

	// EnterSigning is called when entering the signing production.
	EnterSigning(c *SigningContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterTimeunits_declaration is called when entering the timeunits_declaration production.
	EnterTimeunits_declaration(c *Timeunits_declarationContext)

	// EnterLifetime is called when entering the lifetime production.
	EnterLifetime(c *LifetimeContext)

	// EnterPort_direction is called when entering the port_direction production.
	EnterPort_direction(c *Port_directionContext)

	// EnterAlways_keyword is called when entering the always_keyword production.
	EnterAlways_keyword(c *Always_keywordContext)

	// EnterJoin_keyword is called when entering the join_keyword production.
	EnterJoin_keyword(c *Join_keywordContext)

	// EnterUnique_priority is called when entering the unique_priority production.
	EnterUnique_priority(c *Unique_priorityContext)

	// EnterDrive_strength is called when entering the drive_strength production.
	EnterDrive_strength(c *Drive_strengthContext)

	// EnterStrength0 is called when entering the strength0 production.
	EnterStrength0(c *Strength0Context)

	// EnterStrength1 is called when entering the strength1 production.
	EnterStrength1(c *Strength1Context)

	// EnterCharge_strength is called when entering the charge_strength production.
	EnterCharge_strength(c *Charge_strengthContext)

	// EnterSequence_lvar_port_direction is called when entering the sequence_lvar_port_direction production.
	EnterSequence_lvar_port_direction(c *Sequence_lvar_port_directionContext)

	// EnterBins_keyword is called when entering the bins_keyword production.
	EnterBins_keyword(c *Bins_keywordContext)

	// EnterClass_item_qualifier is called when entering the class_item_qualifier production.
	EnterClass_item_qualifier(c *Class_item_qualifierContext)

	// EnterRandom_qualifier is called when entering the random_qualifier production.
	EnterRandom_qualifier(c *Random_qualifierContext)

	// EnterProperty_qualifier is called when entering the property_qualifier production.
	EnterProperty_qualifier(c *Property_qualifierContext)

	// EnterMethod_qualifier is called when entering the method_qualifier production.
	EnterMethod_qualifier(c *Method_qualifierContext)

	// EnterConstraint_prototype_qualifier is called when entering the constraint_prototype_qualifier production.
	EnterConstraint_prototype_qualifier(c *Constraint_prototype_qualifierContext)

	// EnterCmos_switchtype is called when entering the cmos_switchtype production.
	EnterCmos_switchtype(c *Cmos_switchtypeContext)

	// EnterEnable_gatetype is called when entering the enable_gatetype production.
	EnterEnable_gatetype(c *Enable_gatetypeContext)

	// EnterMos_switchtype is called when entering the mos_switchtype production.
	EnterMos_switchtype(c *Mos_switchtypeContext)

	// EnterN_input_gatetype is called when entering the n_input_gatetype production.
	EnterN_input_gatetype(c *N_input_gatetypeContext)

	// EnterN_output_gatetype is called when entering the n_output_gatetype production.
	EnterN_output_gatetype(c *N_output_gatetypeContext)

	// EnterPass_en_switchtype is called when entering the pass_en_switchtype production.
	EnterPass_en_switchtype(c *Pass_en_switchtypeContext)

	// EnterPass_switchtype is called when entering the pass_switchtype production.
	EnterPass_switchtype(c *Pass_switchtypeContext)

	// EnterAny_implication is called when entering the any_implication production.
	EnterAny_implication(c *Any_implicationContext)

	// EnterPolarity_operator is called when entering the polarity_operator production.
	EnterPolarity_operator(c *Polarity_operatorContext)

	// EnterTiming_check_event_control is called when entering the timing_check_event_control production.
	EnterTiming_check_event_control(c *Timing_check_event_controlContext)

	// EnterImport_export is called when entering the import_export production.
	EnterImport_export(c *Import_exportContext)

	// EnterArray_method_name is called when entering the array_method_name production.
	EnterArray_method_name(c *Array_method_nameContext)

	// EnterOperator_mul_div_mod is called when entering the operator_mul_div_mod production.
	EnterOperator_mul_div_mod(c *Operator_mul_div_modContext)

	// EnterOperator_plus_minus is called when entering the operator_plus_minus production.
	EnterOperator_plus_minus(c *Operator_plus_minusContext)

	// EnterOperator_shift is called when entering the operator_shift production.
	EnterOperator_shift(c *Operator_shiftContext)

	// EnterOperator_cmp is called when entering the operator_cmp production.
	EnterOperator_cmp(c *Operator_cmpContext)

	// EnterOperator_eq_neq is called when entering the operator_eq_neq production.
	EnterOperator_eq_neq(c *Operator_eq_neqContext)

	// EnterOperator_xor is called when entering the operator_xor production.
	EnterOperator_xor(c *Operator_xorContext)

	// EnterOperator_impl is called when entering the operator_impl production.
	EnterOperator_impl(c *Operator_implContext)

	// EnterUdp_nonansi_declaration is called when entering the udp_nonansi_declaration production.
	EnterUdp_nonansi_declaration(c *Udp_nonansi_declarationContext)

	// EnterUdp_ansi_declaration is called when entering the udp_ansi_declaration production.
	EnterUdp_ansi_declaration(c *Udp_ansi_declarationContext)

	// EnterUdp_declaration is called when entering the udp_declaration production.
	EnterUdp_declaration(c *Udp_declarationContext)

	// EnterUdp_declaration_port_list is called when entering the udp_declaration_port_list production.
	EnterUdp_declaration_port_list(c *Udp_declaration_port_listContext)

	// EnterUdp_port_declaration is called when entering the udp_port_declaration production.
	EnterUdp_port_declaration(c *Udp_port_declarationContext)

	// EnterUdp_output_declaration is called when entering the udp_output_declaration production.
	EnterUdp_output_declaration(c *Udp_output_declarationContext)

	// EnterUdp_input_declaration is called when entering the udp_input_declaration production.
	EnterUdp_input_declaration(c *Udp_input_declarationContext)

	// EnterUdp_reg_declaration is called when entering the udp_reg_declaration production.
	EnterUdp_reg_declaration(c *Udp_reg_declarationContext)

	// EnterUdp_body is called when entering the udp_body production.
	EnterUdp_body(c *Udp_bodyContext)

	// EnterCombinational_body is called when entering the combinational_body production.
	EnterCombinational_body(c *Combinational_bodyContext)

	// EnterCombinational_entry is called when entering the combinational_entry production.
	EnterCombinational_entry(c *Combinational_entryContext)

	// EnterSequential_body is called when entering the sequential_body production.
	EnterSequential_body(c *Sequential_bodyContext)

	// EnterUdp_initial_statement is called when entering the udp_initial_statement production.
	EnterUdp_initial_statement(c *Udp_initial_statementContext)

	// EnterSequential_entry is called when entering the sequential_entry production.
	EnterSequential_entry(c *Sequential_entryContext)

	// EnterSeq_input_list is called when entering the seq_input_list production.
	EnterSeq_input_list(c *Seq_input_listContext)

	// EnterLevel_input_list is called when entering the level_input_list production.
	EnterLevel_input_list(c *Level_input_listContext)

	// EnterEdge_input_list is called when entering the edge_input_list production.
	EnterEdge_input_list(c *Edge_input_listContext)

	// EnterEdge_indicator is called when entering the edge_indicator production.
	EnterEdge_indicator(c *Edge_indicatorContext)

	// EnterCurrent_state is called when entering the current_state production.
	EnterCurrent_state(c *Current_stateContext)

	// EnterNext_state is called when entering the next_state production.
	EnterNext_state(c *Next_stateContext)

	// EnterInterface_declaration is called when entering the interface_declaration production.
	EnterInterface_declaration(c *Interface_declarationContext)

	// EnterInterface_header is called when entering the interface_header production.
	EnterInterface_header(c *Interface_headerContext)

	// EnterInterface_item is called when entering the interface_item production.
	EnterInterface_item(c *Interface_itemContext)

	// EnterModport_declaration is called when entering the modport_declaration production.
	EnterModport_declaration(c *Modport_declarationContext)

	// EnterModport_item is called when entering the modport_item production.
	EnterModport_item(c *Modport_itemContext)

	// EnterModport_ports_declaration is called when entering the modport_ports_declaration production.
	EnterModport_ports_declaration(c *Modport_ports_declarationContext)

	// EnterModport_clocking_declaration is called when entering the modport_clocking_declaration production.
	EnterModport_clocking_declaration(c *Modport_clocking_declarationContext)

	// EnterModport_simple_ports_declaration is called when entering the modport_simple_ports_declaration production.
	EnterModport_simple_ports_declaration(c *Modport_simple_ports_declarationContext)

	// EnterModport_simple_port is called when entering the modport_simple_port production.
	EnterModport_simple_port(c *Modport_simple_portContext)

	// EnterModport_tf_ports_declaration is called when entering the modport_tf_ports_declaration production.
	EnterModport_tf_ports_declaration(c *Modport_tf_ports_declarationContext)

	// EnterModport_tf_port is called when entering the modport_tf_port production.
	EnterModport_tf_port(c *Modport_tf_portContext)

	// EnterStatement_or_null is called when entering the statement_or_null production.
	EnterStatement_or_null(c *Statement_or_nullContext)

	// EnterInitial_construct is called when entering the initial_construct production.
	EnterInitial_construct(c *Initial_constructContext)

	// EnterDefault_clocking_or_dissable_construct is called when entering the default_clocking_or_dissable_construct production.
	EnterDefault_clocking_or_dissable_construct(c *Default_clocking_or_dissable_constructContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStmtItemMain is called when entering the StmtItemMain production.
	EnterStmtItemMain(c *StmtItemMainContext)

	// EnterStmtItemMacro is called when entering the StmtItemMacro production.
	EnterStmtItemMacro(c *StmtItemMacroContext)

	// EnterStmtItemCase is called when entering the StmtItemCase production.
	EnterStmtItemCase(c *StmtItemCaseContext)

	// EnterStmtItemCond is called when entering the StmtItemCond production.
	EnterStmtItemCond(c *StmtItemCondContext)

	// EnterStmtItemSubCall is called when entering the StmtItemSubCall production.
	EnterStmtItemSubCall(c *StmtItemSubCallContext)

	// EnterStmtItemDisable is called when entering the StmtItemDisable production.
	EnterStmtItemDisable(c *StmtItemDisableContext)

	// EnterStmtItemEvent is called when entering the StmtItemEvent production.
	EnterStmtItemEvent(c *StmtItemEventContext)

	// EnterStmtItemLoop is called when entering the StmtItemLoop production.
	EnterStmtItemLoop(c *StmtItemLoopContext)

	// EnterStmtItemJump is called when entering the StmtItemJump production.
	EnterStmtItemJump(c *StmtItemJumpContext)

	// EnterStmtItemPar is called when entering the StmtItemPar production.
	EnterStmtItemPar(c *StmtItemParContext)

	// EnterStmtItemProcTime is called when entering the StmtItemProcTime production.
	EnterStmtItemProcTime(c *StmtItemProcTimeContext)

	// EnterStmtItemSeq is called when entering the StmtItemSeq production.
	EnterStmtItemSeq(c *StmtItemSeqContext)

	// EnterStmtItemWait is called when entering the StmtItemWait production.
	EnterStmtItemWait(c *StmtItemWaitContext)

	// EnterStmtItemProcAssert is called when entering the StmtItemProcAssert production.
	EnterStmtItemProcAssert(c *StmtItemProcAssertContext)

	// EnterStmtItemRandseq is called when entering the StmtItemRandseq production.
	EnterStmtItemRandseq(c *StmtItemRandseqContext)

	// EnterStmtItemRandcase is called when entering the StmtItemRandcase production.
	EnterStmtItemRandcase(c *StmtItemRandcaseContext)

	// EnterStmtItemExpect is called when entering the StmtItemExpect production.
	EnterStmtItemExpect(c *StmtItemExpectContext)

	// EnterMacro_statement is called when entering the macro_statement production.
	EnterMacro_statement(c *Macro_statementContext)

	// EnterCycle_delay is called when entering the cycle_delay production.
	EnterCycle_delay(c *Cycle_delayContext)

	// EnterClocking_drive is called when entering the clocking_drive production.
	EnterClocking_drive(c *Clocking_driveContext)

	// EnterClockvar_expression is called when entering the clockvar_expression production.
	EnterClockvar_expression(c *Clockvar_expressionContext)

	// EnterFinal_construct is called when entering the final_construct production.
	EnterFinal_construct(c *Final_constructContext)

	// EnterBlocking_assignment is called when entering the blocking_assignment production.
	EnterBlocking_assignment(c *Blocking_assignmentContext)

	// EnterProcedural_timing_control_statement is called when entering the procedural_timing_control_statement production.
	EnterProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// EnterProcedural_timing_control is called when entering the procedural_timing_control production.
	EnterProcedural_timing_control(c *Procedural_timing_controlContext)

	// EnterEvent_control is called when entering the event_control production.
	EnterEvent_control(c *Event_controlContext)

	// EnterDelay_or_event_control is called when entering the delay_or_event_control production.
	EnterDelay_or_event_control(c *Delay_or_event_controlContext)

	// EnterDelay3 is called when entering the delay3 production.
	EnterDelay3(c *Delay3Context)

	// EnterDelay2 is called when entering the delay2 production.
	EnterDelay2(c *Delay2Context)

	// EnterDelay_value is called when entering the delay_value production.
	EnterDelay_value(c *Delay_valueContext)

	// EnterDelay_control is called when entering the delay_control production.
	EnterDelay_control(c *Delay_controlContext)

	// EnterNonblocking_assignment is called when entering the nonblocking_assignment production.
	EnterNonblocking_assignment(c *Nonblocking_assignmentContext)

	// EnterProcedural_continuous_assignment is called when entering the procedural_continuous_assignment production.
	EnterProcedural_continuous_assignment(c *Procedural_continuous_assignmentContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterAction_block is called when entering the action_block production.
	EnterAction_block(c *Action_blockContext)

	// EnterSeq_block is called when entering the seq_block production.
	EnterSeq_block(c *Seq_blockContext)

	// EnterPar_block is called when entering the par_block production.
	EnterPar_block(c *Par_blockContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_keyword is called when entering the case_keyword production.
	EnterCase_keyword(c *Case_keywordContext)

	// EnterCase_item is called when entering the case_item production.
	EnterCase_item(c *Case_itemContext)

	// EnterCase_pattern_item is called when entering the case_pattern_item production.
	EnterCase_pattern_item(c *Case_pattern_itemContext)

	// EnterCase_inside_item is called when entering the case_inside_item production.
	EnterCase_inside_item(c *Case_inside_itemContext)

	// EnterRandcase_statement is called when entering the randcase_statement production.
	EnterRandcase_statement(c *Randcase_statementContext)

	// EnterRandcase_item is called when entering the randcase_item production.
	EnterRandcase_item(c *Randcase_itemContext)

	// EnterCond_predicate is called when entering the cond_predicate production.
	EnterCond_predicate(c *Cond_predicateContext)

	// EnterConditional_statement is called when entering the conditional_statement production.
	EnterConditional_statement(c *Conditional_statementContext)

	// EnterSubroutine_call_statement is called when entering the subroutine_call_statement production.
	EnterSubroutine_call_statement(c *Subroutine_call_statementContext)

	// EnterDisable_statement is called when entering the disable_statement production.
	EnterDisable_statement(c *Disable_statementContext)

	// EnterEvent_trigger is called when entering the event_trigger production.
	EnterEvent_trigger(c *Event_triggerContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterList_of_variable_assignments is called when entering the list_of_variable_assignments production.
	EnterList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// EnterFor_initialization is called when entering the for_initialization production.
	EnterFor_initialization(c *For_initializationContext)

	// EnterFor_variable_declaration_var_assign is called when entering the for_variable_declaration_var_assign production.
	EnterFor_variable_declaration_var_assign(c *For_variable_declaration_var_assignContext)

	// EnterFor_variable_declaration is called when entering the for_variable_declaration production.
	EnterFor_variable_declaration(c *For_variable_declarationContext)

	// EnterFor_step is called when entering the for_step production.
	EnterFor_step(c *For_stepContext)

	// EnterLoop_variables is called when entering the loop_variables production.
	EnterLoop_variables(c *Loop_variablesContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterName_of_instance is called when entering the name_of_instance production.
	EnterName_of_instance(c *Name_of_instanceContext)

	// EnterChecker_instantiation is called when entering the checker_instantiation production.
	EnterChecker_instantiation(c *Checker_instantiationContext)

	// EnterList_of_checker_port_connections is called when entering the list_of_checker_port_connections production.
	EnterList_of_checker_port_connections(c *List_of_checker_port_connectionsContext)

	// EnterOrdered_checker_port_connection is called when entering the ordered_checker_port_connection production.
	EnterOrdered_checker_port_connection(c *Ordered_checker_port_connectionContext)

	// EnterNamed_checker_port_connection is called when entering the named_checker_port_connection production.
	EnterNamed_checker_port_connection(c *Named_checker_port_connectionContext)

	// EnterProcedural_assertion_statement is called when entering the procedural_assertion_statement production.
	EnterProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// EnterConcurrent_assertion_statement is called when entering the concurrent_assertion_statement production.
	EnterConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// EnterAssertion_item is called when entering the assertion_item production.
	EnterAssertion_item(c *Assertion_itemContext)

	// EnterConcurrent_assertion_item is called when entering the concurrent_assertion_item production.
	EnterConcurrent_assertion_item(c *Concurrent_assertion_itemContext)

	// EnterImmediate_assertion_statement is called when entering the immediate_assertion_statement production.
	EnterImmediate_assertion_statement(c *Immediate_assertion_statementContext)

	// EnterSimple_immediate_assertion_statement is called when entering the simple_immediate_assertion_statement production.
	EnterSimple_immediate_assertion_statement(c *Simple_immediate_assertion_statementContext)

	// EnterSimple_immediate_assert_statement is called when entering the simple_immediate_assert_statement production.
	EnterSimple_immediate_assert_statement(c *Simple_immediate_assert_statementContext)

	// EnterSimple_immediate_assume_statement is called when entering the simple_immediate_assume_statement production.
	EnterSimple_immediate_assume_statement(c *Simple_immediate_assume_statementContext)

	// EnterSimple_immediate_cover_statement is called when entering the simple_immediate_cover_statement production.
	EnterSimple_immediate_cover_statement(c *Simple_immediate_cover_statementContext)

	// EnterDeferred_immediate_assertion_statement is called when entering the deferred_immediate_assertion_statement production.
	EnterDeferred_immediate_assertion_statement(c *Deferred_immediate_assertion_statementContext)

	// EnterPrimitive_delay is called when entering the primitive_delay production.
	EnterPrimitive_delay(c *Primitive_delayContext)

	// EnterDeferred_immediate_assert_statement is called when entering the deferred_immediate_assert_statement production.
	EnterDeferred_immediate_assert_statement(c *Deferred_immediate_assert_statementContext)

	// EnterDeferred_immediate_assume_statement is called when entering the deferred_immediate_assume_statement production.
	EnterDeferred_immediate_assume_statement(c *Deferred_immediate_assume_statementContext)

	// EnterDeferred_immediate_cover_statement is called when entering the deferred_immediate_cover_statement production.
	EnterDeferred_immediate_cover_statement(c *Deferred_immediate_cover_statementContext)

	// EnterWeight_specification is called when entering the weight_specification production.
	EnterWeight_specification(c *Weight_specificationContext)

	// EnterProduction_item is called when entering the production_item production.
	EnterProduction_item(c *Production_itemContext)

	// EnterRs_code_block is called when entering the rs_code_block production.
	EnterRs_code_block(c *Rs_code_blockContext)

	// EnterRandsequence_statement is called when entering the randsequence_statement production.
	EnterRandsequence_statement(c *Randsequence_statementContext)

	// EnterRs_prod is called when entering the rs_prod production.
	EnterRs_prod(c *Rs_prodContext)

	// EnterRs_if_else is called when entering the rs_if_else production.
	EnterRs_if_else(c *Rs_if_elseContext)

	// EnterRs_repeat is called when entering the rs_repeat production.
	EnterRs_repeat(c *Rs_repeatContext)

	// EnterRs_case is called when entering the rs_case production.
	EnterRs_case(c *Rs_caseContext)

	// EnterRs_case_item is called when entering the rs_case_item production.
	EnterRs_case_item(c *Rs_case_itemContext)

	// EnterRs_rule is called when entering the rs_rule production.
	EnterRs_rule(c *Rs_ruleContext)

	// EnterRs_production_list is called when entering the rs_production_list production.
	EnterRs_production_list(c *Rs_production_listContext)

	// EnterProduction is called when entering the production production.
	EnterProduction(c *ProductionContext)

	// EnterTf_item_declaration is called when entering the tf_item_declaration production.
	EnterTf_item_declaration(c *Tf_item_declarationContext)

	// EnterTf_port_list is called when entering the tf_port_list production.
	EnterTf_port_list(c *Tf_port_listContext)

	// EnterTf_port_item is called when entering the tf_port_item production.
	EnterTf_port_item(c *Tf_port_itemContext)

	// EnterTf_port_direction is called when entering the tf_port_direction production.
	EnterTf_port_direction(c *Tf_port_directionContext)

	// EnterTf_port_declaration is called when entering the tf_port_declaration production.
	EnterTf_port_declaration(c *Tf_port_declarationContext)

	// EnterList_of_tf_variable_identifiers_item is called when entering the list_of_tf_variable_identifiers_item production.
	EnterList_of_tf_variable_identifiers_item(c *List_of_tf_variable_identifiers_itemContext)

	// EnterList_of_tf_variable_identifiers is called when entering the list_of_tf_variable_identifiers production.
	EnterList_of_tf_variable_identifiers(c *List_of_tf_variable_identifiersContext)

	// EnterExpect_property_statement is called when entering the expect_property_statement production.
	EnterExpect_property_statement(c *Expect_property_statementContext)

	// EnterBlock_item_declaration is called when entering the block_item_declaration production.
	EnterBlock_item_declaration(c *Block_item_declarationContext)

	// EnterParam_assignment is called when entering the param_assignment production.
	EnterParam_assignment(c *Param_assignmentContext)

	// EnterType_assignment is called when entering the type_assignment production.
	EnterType_assignment(c *Type_assignmentContext)

	// EnterList_of_type_assignments is called when entering the list_of_type_assignments production.
	EnterList_of_type_assignments(c *List_of_type_assignmentsContext)

	// EnterList_of_param_assignments is called when entering the list_of_param_assignments production.
	EnterList_of_param_assignments(c *List_of_param_assignmentsContext)

	// EnterParameter_declaration_primitive is called when entering the parameter_declaration_primitive production.
	EnterParameter_declaration_primitive(c *Parameter_declaration_primitiveContext)

	// EnterLocal_parameter_declaration is called when entering the local_parameter_declaration production.
	EnterLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterNet_type_declaration is called when entering the net_type_declaration production.
	EnterNet_type_declaration(c *Net_type_declarationContext)

	// EnterLet_declaration is called when entering the let_declaration production.
	EnterLet_declaration(c *Let_declarationContext)

	// EnterLet_port_list is called when entering the let_port_list production.
	EnterLet_port_list(c *Let_port_listContext)

	// EnterLet_port_item is called when entering the let_port_item production.
	EnterLet_port_item(c *Let_port_itemContext)

	// EnterLet_formal_type is called when entering the let_formal_type production.
	EnterLet_formal_type(c *Let_formal_typeContext)

	// EnterPackage_import_declaration is called when entering the package_import_declaration production.
	EnterPackage_import_declaration(c *Package_import_declarationContext)

	// EnterPackage_import_item is called when entering the package_import_item production.
	EnterPackage_import_item(c *Package_import_itemContext)

	// EnterProperty_list_of_arguments is called when entering the property_list_of_arguments production.
	EnterProperty_list_of_arguments(c *Property_list_of_argumentsContext)

	// EnterProperty_actual_arg is called when entering the property_actual_arg production.
	EnterProperty_actual_arg(c *Property_actual_argContext)

	// EnterProperty_formal_type is called when entering the property_formal_type production.
	EnterProperty_formal_type(c *Property_formal_typeContext)

	// EnterSequence_formal_type is called when entering the sequence_formal_type production.
	EnterSequence_formal_type(c *Sequence_formal_typeContext)

	// EnterProperty_instance is called when entering the property_instance production.
	EnterProperty_instance(c *Property_instanceContext)

	// EnterProperty_spec is called when entering the property_spec production.
	EnterProperty_spec(c *Property_specContext)

	// EnterProperty_expr is called when entering the property_expr production.
	EnterProperty_expr(c *Property_exprContext)

	// EnterProperty_case_item is called when entering the property_case_item production.
	EnterProperty_case_item(c *Property_case_itemContext)

	// EnterBit_select is called when entering the bit_select production.
	EnterBit_select(c *Bit_selectContext)

	// EnterIdentifier_with_bit_select is called when entering the identifier_with_bit_select production.
	EnterIdentifier_with_bit_select(c *Identifier_with_bit_selectContext)

	// EnterPackage_or_class_scoped_hier_id_with_select is called when entering the package_or_class_scoped_hier_id_with_select production.
	EnterPackage_or_class_scoped_hier_id_with_select(c *Package_or_class_scoped_hier_id_with_selectContext)

	// EnterPackage_or_class_scoped_path_item is called when entering the package_or_class_scoped_path_item production.
	EnterPackage_or_class_scoped_path_item(c *Package_or_class_scoped_path_itemContext)

	// EnterPackage_or_class_scoped_path is called when entering the package_or_class_scoped_path production.
	EnterPackage_or_class_scoped_path(c *Package_or_class_scoped_pathContext)

	// EnterHierarchical_identifier is called when entering the hierarchical_identifier production.
	EnterHierarchical_identifier(c *Hierarchical_identifierContext)

	// EnterPackage_or_class_scoped_id is called when entering the package_or_class_scoped_id production.
	EnterPackage_or_class_scoped_id(c *Package_or_class_scoped_idContext)

	// EnterSelect_ is called when entering the select_ production.
	EnterSelect_(c *Select_Context)

	// EnterEvent_expression_item is called when entering the event_expression_item production.
	EnterEvent_expression_item(c *Event_expression_itemContext)

	// EnterEvent_expression is called when entering the event_expression production.
	EnterEvent_expression(c *Event_expressionContext)

	// EnterBoolean_abbrev is called when entering the boolean_abbrev production.
	EnterBoolean_abbrev(c *Boolean_abbrevContext)

	// EnterSequence_abbrev is called when entering the sequence_abbrev production.
	EnterSequence_abbrev(c *Sequence_abbrevContext)

	// EnterConsecutive_repetition is called when entering the consecutive_repetition production.
	EnterConsecutive_repetition(c *Consecutive_repetitionContext)

	// EnterNon_consecutive_repetition is called when entering the non_consecutive_repetition production.
	EnterNon_consecutive_repetition(c *Non_consecutive_repetitionContext)

	// EnterGoto_repetition is called when entering the goto_repetition production.
	EnterGoto_repetition(c *Goto_repetitionContext)

	// EnterCycle_delay_const_range_expression is called when entering the cycle_delay_const_range_expression production.
	EnterCycle_delay_const_range_expression(c *Cycle_delay_const_range_expressionContext)

	// EnterSequence_instance is called when entering the sequence_instance production.
	EnterSequence_instance(c *Sequence_instanceContext)

	// EnterSequence_expr is called when entering the sequence_expr production.
	EnterSequence_expr(c *Sequence_exprContext)

	// EnterSequence_match_item is called when entering the sequence_match_item production.
	EnterSequence_match_item(c *Sequence_match_itemContext)

	// EnterOperator_assignment is called when entering the operator_assignment production.
	EnterOperator_assignment(c *Operator_assignmentContext)

	// EnterSequence_actual_arg is called when entering the sequence_actual_arg production.
	EnterSequence_actual_arg(c *Sequence_actual_argContext)

	// EnterDist_weight is called when entering the dist_weight production.
	EnterDist_weight(c *Dist_weightContext)

	// EnterClocking_declaration is called when entering the clocking_declaration production.
	EnterClocking_declaration(c *Clocking_declarationContext)

	// EnterClocking_item is called when entering the clocking_item production.
	EnterClocking_item(c *Clocking_itemContext)

	// EnterList_of_clocking_decl_assign is called when entering the list_of_clocking_decl_assign production.
	EnterList_of_clocking_decl_assign(c *List_of_clocking_decl_assignContext)

	// EnterClocking_decl_assign is called when entering the clocking_decl_assign production.
	EnterClocking_decl_assign(c *Clocking_decl_assignContext)

	// EnterDefault_skew is called when entering the default_skew production.
	EnterDefault_skew(c *Default_skewContext)

	// EnterClocking_direction is called when entering the clocking_direction production.
	EnterClocking_direction(c *Clocking_directionContext)

	// EnterClocking_skew is called when entering the clocking_skew production.
	EnterClocking_skew(c *Clocking_skewContext)

	// EnterClocking_event is called when entering the clocking_event production.
	EnterClocking_event(c *Clocking_eventContext)

	// EnterCycle_delay_range is called when entering the cycle_delay_range production.
	EnterCycle_delay_range(c *Cycle_delay_rangeContext)

	// EnterExpression_or_dist is called when entering the expression_or_dist production.
	EnterExpression_or_dist(c *Expression_or_distContext)

	// EnterCovergroup_declaration is called when entering the covergroup_declaration production.
	EnterCovergroup_declaration(c *Covergroup_declarationContext)

	// EnterCover_cross is called when entering the cover_cross production.
	EnterCover_cross(c *Cover_crossContext)

	// EnterIdentifier_list_2plus is called when entering the identifier_list_2plus production.
	EnterIdentifier_list_2plus(c *Identifier_list_2plusContext)

	// EnterCross_body is called when entering the cross_body production.
	EnterCross_body(c *Cross_bodyContext)

	// EnterCross_body_item is called when entering the cross_body_item production.
	EnterCross_body_item(c *Cross_body_itemContext)

	// EnterBins_selection_or_option is called when entering the bins_selection_or_option production.
	EnterBins_selection_or_option(c *Bins_selection_or_optionContext)

	// EnterBins_selection is called when entering the bins_selection production.
	EnterBins_selection(c *Bins_selectionContext)

	// EnterSelect_expression is called when entering the select_expression production.
	EnterSelect_expression(c *Select_expressionContext)

	// EnterSelect_condition is called when entering the select_condition production.
	EnterSelect_condition(c *Select_conditionContext)

	// EnterBins_expression is called when entering the bins_expression production.
	EnterBins_expression(c *Bins_expressionContext)

	// EnterCovergroup_range_list is called when entering the covergroup_range_list production.
	EnterCovergroup_range_list(c *Covergroup_range_listContext)

	// EnterCovergroup_value_range is called when entering the covergroup_value_range production.
	EnterCovergroup_value_range(c *Covergroup_value_rangeContext)

	// EnterCovergroup_expression is called when entering the covergroup_expression production.
	EnterCovergroup_expression(c *Covergroup_expressionContext)

	// EnterCoverage_spec_or_option is called when entering the coverage_spec_or_option production.
	EnterCoverage_spec_or_option(c *Coverage_spec_or_optionContext)

	// EnterCoverage_option is called when entering the coverage_option production.
	EnterCoverage_option(c *Coverage_optionContext)

	// EnterCoverage_spec is called when entering the coverage_spec production.
	EnterCoverage_spec(c *Coverage_specContext)

	// EnterCover_point is called when entering the cover_point production.
	EnterCover_point(c *Cover_pointContext)

	// EnterBins_or_empty is called when entering the bins_or_empty production.
	EnterBins_or_empty(c *Bins_or_emptyContext)

	// EnterBins_or_options is called when entering the bins_or_options production.
	EnterBins_or_options(c *Bins_or_optionsContext)

	// EnterTrans_list is called when entering the trans_list production.
	EnterTrans_list(c *Trans_listContext)

	// EnterTrans_set is called when entering the trans_set production.
	EnterTrans_set(c *Trans_setContext)

	// EnterTrans_range_list is called when entering the trans_range_list production.
	EnterTrans_range_list(c *Trans_range_listContext)

	// EnterRepeat_range is called when entering the repeat_range production.
	EnterRepeat_range(c *Repeat_rangeContext)

	// EnterCoverage_event is called when entering the coverage_event production.
	EnterCoverage_event(c *Coverage_eventContext)

	// EnterBlock_event_expression is called when entering the block_event_expression production.
	EnterBlock_event_expression(c *Block_event_expressionContext)

	// EnterHierarchical_btf_identifier is called when entering the hierarchical_btf_identifier production.
	EnterHierarchical_btf_identifier(c *Hierarchical_btf_identifierContext)

	// EnterAssertion_variable_declaration is called when entering the assertion_variable_declaration production.
	EnterAssertion_variable_declaration(c *Assertion_variable_declarationContext)

	// EnterDist_item is called when entering the dist_item production.
	EnterDist_item(c *Dist_itemContext)

	// EnterValue_range is called when entering the value_range production.
	EnterValue_range(c *Value_rangeContext)

	// EnterAttribute_instance is called when entering the attribute_instance production.
	EnterAttribute_instance(c *Attribute_instanceContext)

	// EnterAttr_spec is called when entering the attr_spec production.
	EnterAttr_spec(c *Attr_specContext)

	// EnterClass_new is called when entering the class_new production.
	EnterClass_new(c *Class_newContext)

	// EnterParam_expression is called when entering the param_expression production.
	EnterParam_expression(c *Param_expressionContext)

	// EnterConstant_param_expression is called when entering the constant_param_expression production.
	EnterConstant_param_expression(c *Constant_param_expressionContext)

	// EnterUnpacked_dimension is called when entering the unpacked_dimension production.
	EnterUnpacked_dimension(c *Unpacked_dimensionContext)

	// EnterPacked_dimension is called when entering the packed_dimension production.
	EnterPacked_dimension(c *Packed_dimensionContext)

	// EnterVariable_dimension is called when entering the variable_dimension production.
	EnterVariable_dimension(c *Variable_dimensionContext)

	// EnterStruct_union is called when entering the struct_union production.
	EnterStruct_union(c *Struct_unionContext)

	// EnterEnum_base_type is called when entering the enum_base_type production.
	EnterEnum_base_type(c *Enum_base_typeContext)

	// EnterData_type_primitive is called when entering the data_type_primitive production.
	EnterData_type_primitive(c *Data_type_primitiveContext)

	// EnterData_type_usual is called when entering the data_type_usual production.
	EnterData_type_usual(c *Data_type_usualContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterData_type_or_implicit is called when entering the data_type_or_implicit production.
	EnterData_type_or_implicit(c *Data_type_or_implicitContext)

	// EnterImplicit_data_type is called when entering the implicit_data_type production.
	EnterImplicit_data_type(c *Implicit_data_typeContext)

	// EnterSequence_list_of_arguments_named_item is called when entering the sequence_list_of_arguments_named_item production.
	EnterSequence_list_of_arguments_named_item(c *Sequence_list_of_arguments_named_itemContext)

	// EnterSequence_list_of_arguments is called when entering the sequence_list_of_arguments production.
	EnterSequence_list_of_arguments(c *Sequence_list_of_argumentsContext)

	// EnterList_of_arguments_named_item is called when entering the list_of_arguments_named_item production.
	EnterList_of_arguments_named_item(c *List_of_arguments_named_itemContext)

	// EnterList_of_arguments is called when entering the list_of_arguments production.
	EnterList_of_arguments(c *List_of_argumentsContext)

	// EnterPrimary_literal is called when entering the primary_literal production.
	EnterPrimary_literal(c *Primary_literalContext)

	// EnterType_reference is called when entering the type_reference production.
	EnterType_reference(c *Type_referenceContext)

	// EnterPackage_scope is called when entering the package_scope production.
	EnterPackage_scope(c *Package_scopeContext)

	// EnterPs_identifier is called when entering the ps_identifier production.
	EnterPs_identifier(c *Ps_identifierContext)

	// EnterList_of_parameter_value_assignments is called when entering the list_of_parameter_value_assignments production.
	EnterList_of_parameter_value_assignments(c *List_of_parameter_value_assignmentsContext)

	// EnterParameter_value_assignment is called when entering the parameter_value_assignment production.
	EnterParameter_value_assignment(c *Parameter_value_assignmentContext)

	// EnterClass_type is called when entering the class_type production.
	EnterClass_type(c *Class_typeContext)

	// EnterClass_scope is called when entering the class_scope production.
	EnterClass_scope(c *Class_scopeContext)

	// EnterRange_expression is called when entering the range_expression production.
	EnterRange_expression(c *Range_expressionContext)

	// EnterConstant_range_expression is called when entering the constant_range_expression production.
	EnterConstant_range_expression(c *Constant_range_expressionContext)

	// EnterConstant_mintypmax_expression is called when entering the constant_mintypmax_expression production.
	EnterConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// EnterMintypmax_expression is called when entering the mintypmax_expression production.
	EnterMintypmax_expression(c *Mintypmax_expressionContext)

	// EnterNamed_parameter_assignment is called when entering the named_parameter_assignment production.
	EnterNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// EnterPrimaryLit is called when entering the PrimaryLit production.
	EnterPrimaryLit(c *PrimaryLitContext)

	// EnterPrimaryRandomize is called when entering the PrimaryRandomize production.
	EnterPrimaryRandomize(c *PrimaryRandomizeContext)

	// EnterPrimaryAssig is called when entering the PrimaryAssig production.
	EnterPrimaryAssig(c *PrimaryAssigContext)

	// EnterPrimaryBitSelect is called when entering the PrimaryBitSelect production.
	EnterPrimaryBitSelect(c *PrimaryBitSelectContext)

	// EnterPrimaryTfCall is called when entering the PrimaryTfCall production.
	EnterPrimaryTfCall(c *PrimaryTfCallContext)

	// EnterPrimaryTypeRef is called when entering the PrimaryTypeRef production.
	EnterPrimaryTypeRef(c *PrimaryTypeRefContext)

	// EnterPrimaryCallArrayMethodNoArgs is called when entering the PrimaryCallArrayMethodNoArgs production.
	EnterPrimaryCallArrayMethodNoArgs(c *PrimaryCallArrayMethodNoArgsContext)

	// EnterPrimaryCast is called when entering the PrimaryCast production.
	EnterPrimaryCast(c *PrimaryCastContext)

	// EnterPrimaryPar is called when entering the PrimaryPar production.
	EnterPrimaryPar(c *PrimaryParContext)

	// EnterPrimaryCall is called when entering the PrimaryCall production.
	EnterPrimaryCall(c *PrimaryCallContext)

	// EnterPrimaryRandomize2 is called when entering the PrimaryRandomize2 production.
	EnterPrimaryRandomize2(c *PrimaryRandomize2Context)

	// EnterPrimaryDot is called when entering the PrimaryDot production.
	EnterPrimaryDot(c *PrimaryDotContext)

	// EnterPrimaryStreaming_concatenation is called when entering the PrimaryStreaming_concatenation production.
	EnterPrimaryStreaming_concatenation(c *PrimaryStreaming_concatenationContext)

	// EnterPrimaryPath is called when entering the PrimaryPath production.
	EnterPrimaryPath(c *PrimaryPathContext)

	// EnterPrimaryRange is called when entering the PrimaryRange production.
	EnterPrimaryRange(c *PrimaryRangeContext)

	// EnterPrimaryCallWith is called when entering the PrimaryCallWith production.
	EnterPrimaryCallWith(c *PrimaryCallWithContext)

	// EnterPrimaryConcat is called when entering the PrimaryConcat production.
	EnterPrimaryConcat(c *PrimaryConcatContext)

	// EnterPrimaryCast2 is called when entering the PrimaryCast2 production.
	EnterPrimaryCast2(c *PrimaryCast2Context)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterInc_or_dec_expressionPre is called when entering the Inc_or_dec_expressionPre production.
	EnterInc_or_dec_expressionPre(c *Inc_or_dec_expressionPreContext)

	// EnterInc_or_dec_expressionPost is called when entering the Inc_or_dec_expressionPost production.
	EnterInc_or_dec_expressionPost(c *Inc_or_dec_expressionPostContext)

	// EnterExpressionPrimary is called when entering the ExpressionPrimary production.
	EnterExpressionPrimary(c *ExpressionPrimaryContext)

	// EnterExpressionInside is called when entering the ExpressionInside production.
	EnterExpressionInside(c *ExpressionInsideContext)

	// EnterExpressionBinOpAnd is called when entering the ExpressionBinOpAnd production.
	EnterExpressionBinOpAnd(c *ExpressionBinOpAndContext)

	// EnterExpressionBinOpPower is called when entering the ExpressionBinOpPower production.
	EnterExpressionBinOpPower(c *ExpressionBinOpPowerContext)

	// EnterExpressionBinOpImpl is called when entering the ExpressionBinOpImpl production.
	EnterExpressionBinOpImpl(c *ExpressionBinOpImplContext)

	// EnterExpressionBinOpEq is called when entering the ExpressionBinOpEq production.
	EnterExpressionBinOpEq(c *ExpressionBinOpEqContext)

	// EnterExpressionTernary is called when entering the ExpressionTernary production.
	EnterExpressionTernary(c *ExpressionTernaryContext)

	// EnterExpressionTaggedId is called when entering the ExpressionTaggedId production.
	EnterExpressionTaggedId(c *ExpressionTaggedIdContext)

	// EnterExpressionUnary is called when entering the ExpressionUnary production.
	EnterExpressionUnary(c *ExpressionUnaryContext)

	// EnterExpressionAssign is called when entering the ExpressionAssign production.
	EnterExpressionAssign(c *ExpressionAssignContext)

	// EnterExpressionIncDec is called when entering the ExpressionIncDec production.
	EnterExpressionIncDec(c *ExpressionIncDecContext)

	// EnterExpressionBinOpMul is called when entering the ExpressionBinOpMul production.
	EnterExpressionBinOpMul(c *ExpressionBinOpMulContext)

	// EnterExpressionBinOpShift is called when entering the ExpressionBinOpShift production.
	EnterExpressionBinOpShift(c *ExpressionBinOpShiftContext)

	// EnterExpressionBinOpCmp is called when entering the ExpressionBinOpCmp production.
	EnterExpressionBinOpCmp(c *ExpressionBinOpCmpContext)

	// EnterExpressionBinOpBitAnd is called when entering the ExpressionBinOpBitAnd production.
	EnterExpressionBinOpBitAnd(c *ExpressionBinOpBitAndContext)

	// EnterExpressionBinOpAdd is called when entering the ExpressionBinOpAdd production.
	EnterExpressionBinOpAdd(c *ExpressionBinOpAddContext)

	// EnterExpressionBinOpBitXor is called when entering the ExpressionBinOpBitXor production.
	EnterExpressionBinOpBitXor(c *ExpressionBinOpBitXorContext)

	// EnterExpressionBinOpBitOr is called when entering the ExpressionBinOpBitOr production.
	EnterExpressionBinOpBitOr(c *ExpressionBinOpBitOrContext)

	// EnterExpressionBinOpOr is called when entering the ExpressionBinOpOr production.
	EnterExpressionBinOpOr(c *ExpressionBinOpOrContext)

	// EnterExpressionTripleAnd is called when entering the ExpressionTripleAnd production.
	EnterExpressionTripleAnd(c *ExpressionTripleAndContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterDynamic_array_new is called when entering the dynamic_array_new production.
	EnterDynamic_array_new(c *Dynamic_array_newContext)

	// EnterConst_or_range_expression is called when entering the const_or_range_expression production.
	EnterConst_or_range_expression(c *Const_or_range_expressionContext)

	// EnterVariable_decl_assignment is called when entering the variable_decl_assignment production.
	EnterVariable_decl_assignment(c *Variable_decl_assignmentContext)

	// EnterAssignment_pattern_variable_lvalue is called when entering the assignment_pattern_variable_lvalue production.
	EnterAssignment_pattern_variable_lvalue(c *Assignment_pattern_variable_lvalueContext)

	// EnterStream_operator is called when entering the stream_operator production.
	EnterStream_operator(c *Stream_operatorContext)

	// EnterSlice_size is called when entering the slice_size production.
	EnterSlice_size(c *Slice_sizeContext)

	// EnterStreaming_concatenation is called when entering the streaming_concatenation production.
	EnterStreaming_concatenation(c *Streaming_concatenationContext)

	// EnterStream_concatenation is called when entering the stream_concatenation production.
	EnterStream_concatenation(c *Stream_concatenationContext)

	// EnterStream_expression is called when entering the stream_expression production.
	EnterStream_expression(c *Stream_expressionContext)

	// EnterArray_range_expression is called when entering the array_range_expression production.
	EnterArray_range_expression(c *Array_range_expressionContext)

	// EnterOpen_range_list is called when entering the open_range_list production.
	EnterOpen_range_list(c *Open_range_listContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterAssignment_pattern is called when entering the assignment_pattern production.
	EnterAssignment_pattern(c *Assignment_patternContext)

	// EnterStructure_pattern_key is called when entering the structure_pattern_key production.
	EnterStructure_pattern_key(c *Structure_pattern_keyContext)

	// EnterArray_pattern_key is called when entering the array_pattern_key production.
	EnterArray_pattern_key(c *Array_pattern_keyContext)

	// EnterAssignment_pattern_key is called when entering the assignment_pattern_key production.
	EnterAssignment_pattern_key(c *Assignment_pattern_keyContext)

	// EnterStruct_union_member is called when entering the struct_union_member production.
	EnterStruct_union_member(c *Struct_union_memberContext)

	// EnterData_type_or_void is called when entering the data_type_or_void production.
	EnterData_type_or_void(c *Data_type_or_voidContext)

	// EnterEnum_name_declaration is called when entering the enum_name_declaration production.
	EnterEnum_name_declaration(c *Enum_name_declarationContext)

	// EnterAssignment_pattern_expression is called when entering the assignment_pattern_expression production.
	EnterAssignment_pattern_expression(c *Assignment_pattern_expressionContext)

	// EnterAssignment_pattern_expression_type is called when entering the assignment_pattern_expression_type production.
	EnterAssignment_pattern_expression_type(c *Assignment_pattern_expression_typeContext)

	// EnterNet_lvalue is called when entering the net_lvalue production.
	EnterNet_lvalue(c *Net_lvalueContext)

	// EnterVarLConcat is called when entering the VarLConcat production.
	EnterVarLConcat(c *VarLConcatContext)

	// EnterVarLPath is called when entering the VarLPath production.
	EnterVarLPath(c *VarLPathContext)

	// EnterVarLAssign is called when entering the VarLAssign production.
	EnterVarLAssign(c *VarLAssignContext)

	// EnterVarLStreamConcat is called when entering the VarLStreamConcat production.
	EnterVarLStreamConcat(c *VarLStreamConcatContext)

	// EnterSolve_before_list is called when entering the solve_before_list production.
	EnterSolve_before_list(c *Solve_before_listContext)

	// EnterConstraint_block_item is called when entering the constraint_block_item production.
	EnterConstraint_block_item(c *Constraint_block_itemContext)

	// EnterConstraint_expression is called when entering the constraint_expression production.
	EnterConstraint_expression(c *Constraint_expressionContext)

	// EnterUniqueness_constraint is called when entering the uniqueness_constraint production.
	EnterUniqueness_constraint(c *Uniqueness_constraintContext)

	// EnterConstraint_set is called when entering the constraint_set production.
	EnterConstraint_set(c *Constraint_setContext)

	// EnterRandomize_call is called when entering the randomize_call production.
	EnterRandomize_call(c *Randomize_callContext)

	// EnterModule_header_common is called when entering the module_header_common production.
	EnterModule_header_common(c *Module_header_commonContext)

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterModule_keyword is called when entering the module_keyword production.
	EnterModule_keyword(c *Module_keywordContext)

	// EnterNet_port_type is called when entering the net_port_type production.
	EnterNet_port_type(c *Net_port_typeContext)

	// EnterVar_data_type is called when entering the var_data_type production.
	EnterVar_data_type(c *Var_data_typeContext)

	// EnterNet_or_var_data_type is called when entering the net_or_var_data_type production.
	EnterNet_or_var_data_type(c *Net_or_var_data_typeContext)

	// EnterList_of_defparam_assignments is called when entering the list_of_defparam_assignments production.
	EnterList_of_defparam_assignments(c *List_of_defparam_assignmentsContext)

	// EnterList_of_net_decl_assignments is called when entering the list_of_net_decl_assignments production.
	EnterList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// EnterList_of_specparam_assignments is called when entering the list_of_specparam_assignments production.
	EnterList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// EnterList_of_variable_decl_assignments is called when entering the list_of_variable_decl_assignments production.
	EnterList_of_variable_decl_assignments(c *List_of_variable_decl_assignmentsContext)

	// EnterList_of_variable_identifiers_item is called when entering the list_of_variable_identifiers_item production.
	EnterList_of_variable_identifiers_item(c *List_of_variable_identifiers_itemContext)

	// EnterList_of_variable_identifiers is called when entering the list_of_variable_identifiers production.
	EnterList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// EnterList_of_variable_port_identifiers is called when entering the list_of_variable_port_identifiers production.
	EnterList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// EnterDefparam_assignment is called when entering the defparam_assignment production.
	EnterDefparam_assignment(c *Defparam_assignmentContext)

	// EnterNet_decl_assignment is called when entering the net_decl_assignment production.
	EnterNet_decl_assignment(c *Net_decl_assignmentContext)

	// EnterSpecparam_assignment is called when entering the specparam_assignment production.
	EnterSpecparam_assignment(c *Specparam_assignmentContext)

	// EnterError_limit_value is called when entering the error_limit_value production.
	EnterError_limit_value(c *Error_limit_valueContext)

	// EnterReject_limit_value is called when entering the reject_limit_value production.
	EnterReject_limit_value(c *Reject_limit_valueContext)

	// EnterPulse_control_specparam is called when entering the pulse_control_specparam production.
	EnterPulse_control_specparam(c *Pulse_control_specparamContext)

	// EnterIdentifier_doted_index_at_end is called when entering the identifier_doted_index_at_end production.
	EnterIdentifier_doted_index_at_end(c *Identifier_doted_index_at_endContext)

	// EnterSpecify_terminal_descriptor is called when entering the specify_terminal_descriptor production.
	EnterSpecify_terminal_descriptor(c *Specify_terminal_descriptorContext)

	// EnterSpecify_input_terminal_descriptor is called when entering the specify_input_terminal_descriptor production.
	EnterSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// EnterSpecify_output_terminal_descriptor is called when entering the specify_output_terminal_descriptor production.
	EnterSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// EnterSpecify_item is called when entering the specify_item production.
	EnterSpecify_item(c *Specify_itemContext)

	// EnterPulsestyle_declaration is called when entering the pulsestyle_declaration production.
	EnterPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// EnterShowcancelled_declaration is called when entering the showcancelled_declaration production.
	EnterShowcancelled_declaration(c *Showcancelled_declarationContext)

	// EnterPath_declaration is called when entering the path_declaration production.
	EnterPath_declaration(c *Path_declarationContext)

	// EnterSimple_path_declaration is called when entering the simple_path_declaration production.
	EnterSimple_path_declaration(c *Simple_path_declarationContext)

	// EnterPath_delay_value is called when entering the path_delay_value production.
	EnterPath_delay_value(c *Path_delay_valueContext)

	// EnterList_of_path_outputs is called when entering the list_of_path_outputs production.
	EnterList_of_path_outputs(c *List_of_path_outputsContext)

	// EnterList_of_path_inputs is called when entering the list_of_path_inputs production.
	EnterList_of_path_inputs(c *List_of_path_inputsContext)

	// EnterList_of_paths is called when entering the list_of_paths production.
	EnterList_of_paths(c *List_of_pathsContext)

	// EnterList_of_path_delay_expressions is called when entering the list_of_path_delay_expressions production.
	EnterList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// EnterT_path_delay_expression is called when entering the t_path_delay_expression production.
	EnterT_path_delay_expression(c *T_path_delay_expressionContext)

	// EnterTrise_path_delay_expression is called when entering the trise_path_delay_expression production.
	EnterTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// EnterTfall_path_delay_expression is called when entering the tfall_path_delay_expression production.
	EnterTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// EnterTz_path_delay_expression is called when entering the tz_path_delay_expression production.
	EnterTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// EnterT01_path_delay_expression is called when entering the t01_path_delay_expression production.
	EnterT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// EnterT10_path_delay_expression is called when entering the t10_path_delay_expression production.
	EnterT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// EnterT0z_path_delay_expression is called when entering the t0z_path_delay_expression production.
	EnterT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// EnterTz1_path_delay_expression is called when entering the tz1_path_delay_expression production.
	EnterTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// EnterT1z_path_delay_expression is called when entering the t1z_path_delay_expression production.
	EnterT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// EnterTz0_path_delay_expression is called when entering the tz0_path_delay_expression production.
	EnterTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// EnterT0x_path_delay_expression is called when entering the t0x_path_delay_expression production.
	EnterT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// EnterTx1_path_delay_expression is called when entering the tx1_path_delay_expression production.
	EnterTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// EnterT1x_path_delay_expression is called when entering the t1x_path_delay_expression production.
	EnterT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// EnterTx0_path_delay_expression is called when entering the tx0_path_delay_expression production.
	EnterTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// EnterTxz_path_delay_expression is called when entering the txz_path_delay_expression production.
	EnterTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// EnterTzx_path_delay_expression is called when entering the tzx_path_delay_expression production.
	EnterTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// EnterParallel_path_description is called when entering the parallel_path_description production.
	EnterParallel_path_description(c *Parallel_path_descriptionContext)

	// EnterFull_path_description is called when entering the full_path_description production.
	EnterFull_path_description(c *Full_path_descriptionContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterSpecparam_declaration is called when entering the specparam_declaration production.
	EnterSpecparam_declaration(c *Specparam_declarationContext)

	// EnterEdge_sensitive_path_declaration is called when entering the edge_sensitive_path_declaration production.
	EnterEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// EnterParallel_edge_sensitive_path_description is called when entering the parallel_edge_sensitive_path_description production.
	EnterParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// EnterFull_edge_sensitive_path_description is called when entering the full_edge_sensitive_path_description production.
	EnterFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// EnterData_source_expression is called when entering the data_source_expression production.
	EnterData_source_expression(c *Data_source_expressionContext)

	// EnterData_declaration is called when entering the data_declaration production.
	EnterData_declaration(c *Data_declarationContext)

	// EnterModule_path_expression is called when entering the module_path_expression production.
	EnterModule_path_expression(c *Module_path_expressionContext)

	// EnterState_dependent_path_declaration is called when entering the state_dependent_path_declaration production.
	EnterState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// EnterPackage_export_declaration is called when entering the package_export_declaration production.
	EnterPackage_export_declaration(c *Package_export_declarationContext)

	// EnterGenvar_declaration is called when entering the genvar_declaration production.
	EnterGenvar_declaration(c *Genvar_declarationContext)

	// EnterNet_declaration is called when entering the net_declaration production.
	EnterNet_declaration(c *Net_declarationContext)

	// EnterParameter_port_list is called when entering the parameter_port_list production.
	EnterParameter_port_list(c *Parameter_port_listContext)

	// EnterParamPortType is called when entering the ParamPortType production.
	EnterParamPortType(c *ParamPortTypeContext)

	// EnterParamSimple is called when entering the ParamSimple production.
	EnterParamSimple(c *ParamSimpleContext)

	// EnterParamLocal is called when entering the ParamLocal production.
	EnterParamLocal(c *ParamLocalContext)

	// EnterParamAssign is called when entering the ParamAssign production.
	EnterParamAssign(c *ParamAssignContext)

	// EnterList_of_port_declarations_ansi_item is called when entering the list_of_port_declarations_ansi_item production.
	EnterList_of_port_declarations_ansi_item(c *List_of_port_declarations_ansi_itemContext)

	// EnterList_of_port_declarations is called when entering the list_of_port_declarations production.
	EnterList_of_port_declarations(c *List_of_port_declarationsContext)

	// EnterNonansi_port_declaration is called when entering the nonansi_port_declaration production.
	EnterNonansi_port_declaration(c *Nonansi_port_declarationContext)

	// EnterNonansi_port is called when entering the nonansi_port production.
	EnterNonansi_port(c *Nonansi_portContext)

	// EnterNonansi_port__expr is called when entering the nonansi_port__expr production.
	EnterNonansi_port__expr(c *Nonansi_port__exprContext)

	// EnterPort_identifier is called when entering the port_identifier production.
	EnterPort_identifier(c *Port_identifierContext)

	// EnterAnsi_port_declaration is called when entering the ansi_port_declaration production.
	EnterAnsi_port_declaration(c *Ansi_port_declarationContext)

	// EnterSystem_timing_check is called when entering the system_timing_check production.
	EnterSystem_timing_check(c *System_timing_checkContext)

	// EnterDolar_setup_timing_check is called when entering the dolar_setup_timing_check production.
	EnterDolar_setup_timing_check(c *Dolar_setup_timing_checkContext)

	// EnterDolar_hold_timing_check is called when entering the dolar_hold_timing_check production.
	EnterDolar_hold_timing_check(c *Dolar_hold_timing_checkContext)

	// EnterDolar_setuphold_timing_check is called when entering the dolar_setuphold_timing_check production.
	EnterDolar_setuphold_timing_check(c *Dolar_setuphold_timing_checkContext)

	// EnterDolar_recovery_timing_check is called when entering the dolar_recovery_timing_check production.
	EnterDolar_recovery_timing_check(c *Dolar_recovery_timing_checkContext)

	// EnterDolar_removal_timing_check is called when entering the dolar_removal_timing_check production.
	EnterDolar_removal_timing_check(c *Dolar_removal_timing_checkContext)

	// EnterDolar_recrem_timing_check is called when entering the dolar_recrem_timing_check production.
	EnterDolar_recrem_timing_check(c *Dolar_recrem_timing_checkContext)

	// EnterDolar_skew_timing_check is called when entering the dolar_skew_timing_check production.
	EnterDolar_skew_timing_check(c *Dolar_skew_timing_checkContext)

	// EnterDolar_timeskew_timing_check is called when entering the dolar_timeskew_timing_check production.
	EnterDolar_timeskew_timing_check(c *Dolar_timeskew_timing_checkContext)

	// EnterDolar_fullskew_timing_check is called when entering the dolar_fullskew_timing_check production.
	EnterDolar_fullskew_timing_check(c *Dolar_fullskew_timing_checkContext)

	// EnterDolar_period_timing_check is called when entering the dolar_period_timing_check production.
	EnterDolar_period_timing_check(c *Dolar_period_timing_checkContext)

	// EnterDolar_width_timing_check is called when entering the dolar_width_timing_check production.
	EnterDolar_width_timing_check(c *Dolar_width_timing_checkContext)

	// EnterDolar_nochange_timing_check is called when entering the dolar_nochange_timing_check production.
	EnterDolar_nochange_timing_check(c *Dolar_nochange_timing_checkContext)

	// EnterTimecheck_condition is called when entering the timecheck_condition production.
	EnterTimecheck_condition(c *Timecheck_conditionContext)

	// EnterControlled_reference_event is called when entering the controlled_reference_event production.
	EnterControlled_reference_event(c *Controlled_reference_eventContext)

	// EnterDelayed_reference is called when entering the delayed_reference production.
	EnterDelayed_reference(c *Delayed_referenceContext)

	// EnterEnd_edge_offset is called when entering the end_edge_offset production.
	EnterEnd_edge_offset(c *End_edge_offsetContext)

	// EnterEvent_based_flag is called when entering the event_based_flag production.
	EnterEvent_based_flag(c *Event_based_flagContext)

	// EnterNotifier is called when entering the notifier production.
	EnterNotifier(c *NotifierContext)

	// EnterRemain_active_flag is called when entering the remain_active_flag production.
	EnterRemain_active_flag(c *Remain_active_flagContext)

	// EnterTimestamp_condition is called when entering the timestamp_condition production.
	EnterTimestamp_condition(c *Timestamp_conditionContext)

	// EnterStart_edge_offset is called when entering the start_edge_offset production.
	EnterStart_edge_offset(c *Start_edge_offsetContext)

	// EnterThreshold is called when entering the threshold production.
	EnterThreshold(c *ThresholdContext)

	// EnterTiming_check_limit is called when entering the timing_check_limit production.
	EnterTiming_check_limit(c *Timing_check_limitContext)

	// EnterTiming_check_event is called when entering the timing_check_event production.
	EnterTiming_check_event(c *Timing_check_eventContext)

	// EnterTiming_check_condition is called when entering the timing_check_condition production.
	EnterTiming_check_condition(c *Timing_check_conditionContext)

	// EnterScalar_timing_check_condition is called when entering the scalar_timing_check_condition production.
	EnterScalar_timing_check_condition(c *Scalar_timing_check_conditionContext)

	// EnterControlled_timing_check_event is called when entering the controlled_timing_check_event production.
	EnterControlled_timing_check_event(c *Controlled_timing_check_eventContext)

	// EnterFunction_data_type_or_implicit is called when entering the function_data_type_or_implicit production.
	EnterFunction_data_type_or_implicit(c *Function_data_type_or_implicitContext)

	// EnterExtern_tf_declaration is called when entering the extern_tf_declaration production.
	EnterExtern_tf_declaration(c *Extern_tf_declarationContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterTask_prototype is called when entering the task_prototype production.
	EnterTask_prototype(c *Task_prototypeContext)

	// EnterFunction_prototype is called when entering the function_prototype production.
	EnterFunction_prototype(c *Function_prototypeContext)

	// EnterDpi_import_export is called when entering the dpi_import_export production.
	EnterDpi_import_export(c *Dpi_import_exportContext)

	// EnterDpi_function_import_property is called when entering the dpi_function_import_property production.
	EnterDpi_function_import_property(c *Dpi_function_import_propertyContext)

	// EnterDpi_task_import_property is called when entering the dpi_task_import_property production.
	EnterDpi_task_import_property(c *Dpi_task_import_propertyContext)

	// EnterTask_and_function_declaration_common is called when entering the task_and_function_declaration_common production.
	EnterTask_and_function_declaration_common(c *Task_and_function_declaration_commonContext)

	// EnterTask_declaration is called when entering the task_declaration production.
	EnterTask_declaration(c *Task_declarationContext)

	// EnterMethod_prototype is called when entering the method_prototype production.
	EnterMethod_prototype(c *Method_prototypeContext)

	// EnterExtern_constraint_declaration is called when entering the extern_constraint_declaration production.
	EnterExtern_constraint_declaration(c *Extern_constraint_declarationContext)

	// EnterConstraint_block is called when entering the constraint_block production.
	EnterConstraint_block(c *Constraint_blockContext)

	// EnterChecker_port_list is called when entering the checker_port_list production.
	EnterChecker_port_list(c *Checker_port_listContext)

	// EnterChecker_port_item is called when entering the checker_port_item production.
	EnterChecker_port_item(c *Checker_port_itemContext)

	// EnterChecker_port_direction is called when entering the checker_port_direction production.
	EnterChecker_port_direction(c *Checker_port_directionContext)

	// EnterChecker_declaration is called when entering the checker_declaration production.
	EnterChecker_declaration(c *Checker_declarationContext)

	// EnterClass_declaration is called when entering the class_declaration production.
	EnterClass_declaration(c *Class_declarationContext)

	// EnterAlways_construct is called when entering the always_construct production.
	EnterAlways_construct(c *Always_constructContext)

	// EnterInterface_class_type is called when entering the interface_class_type production.
	EnterInterface_class_type(c *Interface_class_typeContext)

	// EnterInterface_class_declaration is called when entering the interface_class_declaration production.
	EnterInterface_class_declaration(c *Interface_class_declarationContext)

	// EnterInterface_class_item is called when entering the interface_class_item production.
	EnterInterface_class_item(c *Interface_class_itemContext)

	// EnterInterface_class_method is called when entering the interface_class_method production.
	EnterInterface_class_method(c *Interface_class_methodContext)

	// EnterPackage_declaration is called when entering the package_declaration production.
	EnterPackage_declaration(c *Package_declarationContext)

	// EnterPackage_item is called when entering the package_item production.
	EnterPackage_item(c *Package_itemContext)

	// EnterPackage_item_item is called when entering the package_item_item production.
	EnterPackage_item_item(c *Package_item_itemContext)

	// EnterProgram_declaration is called when entering the program_declaration production.
	EnterProgram_declaration(c *Program_declarationContext)

	// EnterProgram_header is called when entering the program_header production.
	EnterProgram_header(c *Program_headerContext)

	// EnterProgram_item is called when entering the program_item production.
	EnterProgram_item(c *Program_itemContext)

	// EnterNon_port_program_item is called when entering the non_port_program_item production.
	EnterNon_port_program_item(c *Non_port_program_itemContext)

	// EnterAnonymous_program is called when entering the anonymous_program production.
	EnterAnonymous_program(c *Anonymous_programContext)

	// EnterAnonymous_program_item is called when entering the anonymous_program_item production.
	EnterAnonymous_program_item(c *Anonymous_program_itemContext)

	// EnterSequence_declaration is called when entering the sequence_declaration production.
	EnterSequence_declaration(c *Sequence_declarationContext)

	// EnterSequence_port_list is called when entering the sequence_port_list production.
	EnterSequence_port_list(c *Sequence_port_listContext)

	// EnterSequence_port_item is called when entering the sequence_port_item production.
	EnterSequence_port_item(c *Sequence_port_itemContext)

	// EnterProperty_declaration is called when entering the property_declaration production.
	EnterProperty_declaration(c *Property_declarationContext)

	// EnterProperty_port_list is called when entering the property_port_list production.
	EnterProperty_port_list(c *Property_port_listContext)

	// EnterProperty_port_item is called when entering the property_port_item production.
	EnterProperty_port_item(c *Property_port_itemContext)

	// EnterContinuous_assign is called when entering the continuous_assign production.
	EnterContinuous_assign(c *Continuous_assignContext)

	// EnterChecker_or_generate_item is called when entering the checker_or_generate_item production.
	EnterChecker_or_generate_item(c *Checker_or_generate_itemContext)

	// EnterConstraint_prototype is called when entering the constraint_prototype production.
	EnterConstraint_prototype(c *Constraint_prototypeContext)

	// EnterClass_constraint is called when entering the class_constraint production.
	EnterClass_constraint(c *Class_constraintContext)

	// EnterConstraint_declaration is called when entering the constraint_declaration production.
	EnterConstraint_declaration(c *Constraint_declarationContext)

	// EnterClass_constructor_declaration is called when entering the class_constructor_declaration production.
	EnterClass_constructor_declaration(c *Class_constructor_declarationContext)

	// EnterClass_property is called when entering the class_property production.
	EnterClass_property(c *Class_propertyContext)

	// EnterClass_method is called when entering the class_method production.
	EnterClass_method(c *Class_methodContext)

	// EnterClass_constructor_prototype is called when entering the class_constructor_prototype production.
	EnterClass_constructor_prototype(c *Class_constructor_prototypeContext)

	// EnterClass_item is called when entering the class_item production.
	EnterClass_item(c *Class_itemContext)

	// EnterParameter_override is called when entering the parameter_override production.
	EnterParameter_override(c *Parameter_overrideContext)

	// EnterGate_instantiation is called when entering the gate_instantiation production.
	EnterGate_instantiation(c *Gate_instantiationContext)

	// EnterEnable_gate_or_mos_switch_or_cmos_switch_instance is called when entering the enable_gate_or_mos_switch_or_cmos_switch_instance production.
	EnterEnable_gate_or_mos_switch_or_cmos_switch_instance(c *Enable_gate_or_mos_switch_or_cmos_switch_instanceContext)

	// EnterN_input_gate_instance is called when entering the n_input_gate_instance production.
	EnterN_input_gate_instance(c *N_input_gate_instanceContext)

	// EnterN_output_gate_instance is called when entering the n_output_gate_instance production.
	EnterN_output_gate_instance(c *N_output_gate_instanceContext)

	// EnterPass_switch_instance is called when entering the pass_switch_instance production.
	EnterPass_switch_instance(c *Pass_switch_instanceContext)

	// EnterPass_enable_switch_instance is called when entering the pass_enable_switch_instance production.
	EnterPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// EnterPull_gate_instance is called when entering the pull_gate_instance production.
	EnterPull_gate_instance(c *Pull_gate_instanceContext)

	// EnterPulldown_strength is called when entering the pulldown_strength production.
	EnterPulldown_strength(c *Pulldown_strengthContext)

	// EnterPullup_strength is called when entering the pullup_strength production.
	EnterPullup_strength(c *Pullup_strengthContext)

	// EnterEnable_terminal is called when entering the enable_terminal production.
	EnterEnable_terminal(c *Enable_terminalContext)

	// EnterInout_terminal is called when entering the inout_terminal production.
	EnterInout_terminal(c *Inout_terminalContext)

	// EnterInput_terminal is called when entering the input_terminal production.
	EnterInput_terminal(c *Input_terminalContext)

	// EnterOutput_terminal is called when entering the output_terminal production.
	EnterOutput_terminal(c *Output_terminalContext)

	// EnterUdp_instantiation is called when entering the udp_instantiation production.
	EnterUdp_instantiation(c *Udp_instantiationContext)

	// EnterUdp_instance is called when entering the udp_instance production.
	EnterUdp_instance(c *Udp_instanceContext)

	// EnterUdp_instance_body is called when entering the udp_instance_body production.
	EnterUdp_instance_body(c *Udp_instance_bodyContext)

	// EnterModule_or_interface_or_program_or_udp_instantiation is called when entering the module_or_interface_or_program_or_udp_instantiation production.
	EnterModule_or_interface_or_program_or_udp_instantiation(c *Module_or_interface_or_program_or_udp_instantiationContext)

	// EnterHierarchical_instance is called when entering the hierarchical_instance production.
	EnterHierarchical_instance(c *Hierarchical_instanceContext)

	// EnterList_of_port_connections is called when entering the list_of_port_connections production.
	EnterList_of_port_connections(c *List_of_port_connectionsContext)

	// EnterOrdered_port_connection is called when entering the ordered_port_connection production.
	EnterOrdered_port_connection(c *Ordered_port_connectionContext)

	// EnterNamed_port_connection is called when entering the named_port_connection production.
	EnterNamed_port_connection(c *Named_port_connectionContext)

	// EnterBind_directive is called when entering the bind_directive production.
	EnterBind_directive(c *Bind_directiveContext)

	// EnterBind_target_instance is called when entering the bind_target_instance production.
	EnterBind_target_instance(c *Bind_target_instanceContext)

	// EnterBind_target_instance_list is called when entering the bind_target_instance_list production.
	EnterBind_target_instance_list(c *Bind_target_instance_listContext)

	// EnterBind_instantiation is called when entering the bind_instantiation production.
	EnterBind_instantiation(c *Bind_instantiationContext)

	// EnterConfig_declaration is called when entering the config_declaration production.
	EnterConfig_declaration(c *Config_declarationContext)

	// EnterDesign_statement is called when entering the design_statement production.
	EnterDesign_statement(c *Design_statementContext)

	// EnterConfig_rule_statement is called when entering the config_rule_statement production.
	EnterConfig_rule_statement(c *Config_rule_statementContext)

	// EnterInst_clause is called when entering the inst_clause production.
	EnterInst_clause(c *Inst_clauseContext)

	// EnterInst_name is called when entering the inst_name production.
	EnterInst_name(c *Inst_nameContext)

	// EnterCell_clause is called when entering the cell_clause production.
	EnterCell_clause(c *Cell_clauseContext)

	// EnterLiblist_clause is called when entering the liblist_clause production.
	EnterLiblist_clause(c *Liblist_clauseContext)

	// EnterUse_clause is called when entering the use_clause production.
	EnterUse_clause(c *Use_clauseContext)

	// EnterNet_alias is called when entering the net_alias production.
	EnterNet_alias(c *Net_aliasContext)

	// EnterSpecify_block is called when entering the specify_block production.
	EnterSpecify_block(c *Specify_blockContext)

	// EnterGenerate_region is called when entering the generate_region production.
	EnterGenerate_region(c *Generate_regionContext)

	// EnterGenvar_expression is called when entering the genvar_expression production.
	EnterGenvar_expression(c *Genvar_expressionContext)

	// EnterLoop_generate_construct is called when entering the loop_generate_construct production.
	EnterLoop_generate_construct(c *Loop_generate_constructContext)

	// EnterGenvar_initialization is called when entering the genvar_initialization production.
	EnterGenvar_initialization(c *Genvar_initializationContext)

	// EnterGenIterPostfix is called when entering the GenIterPostfix production.
	EnterGenIterPostfix(c *GenIterPostfixContext)

	// EnterGenIterPrefix is called when entering the GenIterPrefix production.
	EnterGenIterPrefix(c *GenIterPrefixContext)

	// EnterConditional_generate_construct is called when entering the conditional_generate_construct production.
	EnterConditional_generate_construct(c *Conditional_generate_constructContext)

	// EnterIf_generate_construct is called when entering the if_generate_construct production.
	EnterIf_generate_construct(c *If_generate_constructContext)

	// EnterCase_generate_construct is called when entering the case_generate_construct production.
	EnterCase_generate_construct(c *Case_generate_constructContext)

	// EnterCase_generate_item is called when entering the case_generate_item production.
	EnterCase_generate_item(c *Case_generate_itemContext)

	// EnterGenerate_begin_end_block is called when entering the generate_begin_end_block production.
	EnterGenerate_begin_end_block(c *Generate_begin_end_blockContext)

	// EnterGenerate_item_item is called when entering the generate_item_item production.
	EnterGenerate_item_item(c *Generate_item_itemContext)

	// EnterGenItemItem is called when entering the GenItemItem production.
	EnterGenItemItem(c *GenItemItemContext)

	// EnterGenItemDataDecl is called when entering the GenItemDataDecl production.
	EnterGenItemDataDecl(c *GenItemDataDeclContext)

	// EnterGenItemGenReg is called when entering the GenItemGenReg production.
	EnterGenItemGenReg(c *GenItemGenRegContext)

	// EnterGenItemGenBlock is called when entering the GenItemGenBlock production.
	EnterGenItemGenBlock(c *GenItemGenBlockContext)

	// EnterProgram_generate_item is called when entering the program_generate_item production.
	EnterProgram_generate_item(c *Program_generate_itemContext)

	// EnterElaboration_system_task is called when entering the elaboration_system_task production.
	EnterElaboration_system_task(c *Elaboration_system_taskContext)

	// EnterCoreItemParamOver is called when entering the CoreItemParamOver production.
	EnterCoreItemParamOver(c *CoreItemParamOverContext)

	// EnterCoreItemGate is called when entering the CoreItemGate production.
	EnterCoreItemGate(c *CoreItemGateContext)

	// EnterCoreItemUdp is called when entering the CoreItemUdp production.
	EnterCoreItemUdp(c *CoreItemUdpContext)

	// EnterCoreItemInstance is called when entering the CoreItemInstance production.
	EnterCoreItemInstance(c *CoreItemInstanceContext)

	// EnterCoreItemParam is called when entering the CoreItemParam production.
	EnterCoreItemParam(c *CoreItemParamContext)

	// EnterCoreItemNet is called when entering the CoreItemNet production.
	EnterCoreItemNet(c *CoreItemNetContext)

	// EnterCoreItemData is called when entering the CoreItemData production.
	EnterCoreItemData(c *CoreItemDataContext)

	// EnterCoreItemTask is called when entering the CoreItemTask production.
	EnterCoreItemTask(c *CoreItemTaskContext)

	// EnterCoreItemFunction is called when entering the CoreItemFunction production.
	EnterCoreItemFunction(c *CoreItemFunctionContext)

	// EnterCoreItemChecker is called when entering the CoreItemChecker production.
	EnterCoreItemChecker(c *CoreItemCheckerContext)

	// EnterCoreItemDPI is called when entering the CoreItemDPI production.
	EnterCoreItemDPI(c *CoreItemDPIContext)

	// EnterCoreItemExtern is called when entering the CoreItemExtern production.
	EnterCoreItemExtern(c *CoreItemExternContext)

	// EnterCoreItemClass is called when entering the CoreItemClass production.
	EnterCoreItemClass(c *CoreItemClassContext)

	// EnterCoreItemIntf is called when entering the CoreItemIntf production.
	EnterCoreItemIntf(c *CoreItemIntfContext)

	// EnterCoreItemClassCons is called when entering the CoreItemClassCons production.
	EnterCoreItemClassCons(c *CoreItemClassConsContext)

	// EnterCoreItemCover is called when entering the CoreItemCover production.
	EnterCoreItemCover(c *CoreItemCoverContext)

	// EnterCoreItemProperty is called when entering the CoreItemProperty production.
	EnterCoreItemProperty(c *CoreItemPropertyContext)

	// EnterCoreItemSeq is called when entering the CoreItemSeq production.
	EnterCoreItemSeq(c *CoreItemSeqContext)

	// EnterCoreItemLet is called when entering the CoreItemLet production.
	EnterCoreItemLet(c *CoreItemLetContext)

	// EnterCoreItemGenVar is called when entering the CoreItemGenVar production.
	EnterCoreItemGenVar(c *CoreItemGenVarContext)

	// EnterCoreItemClock is called when entering the CoreItemClock production.
	EnterCoreItemClock(c *CoreItemClockContext)

	// EnterCoreItemAssert is called when entering the CoreItemAssert production.
	EnterCoreItemAssert(c *CoreItemAssertContext)

	// EnterCoreItemBind is called when entering the CoreItemBind production.
	EnterCoreItemBind(c *CoreItemBindContext)

	// EnterCoreItemContinuous is called when entering the CoreItemContinuous production.
	EnterCoreItemContinuous(c *CoreItemContinuousContext)

	// EnterCoreItemNetAlias is called when entering the CoreItemNetAlias production.
	EnterCoreItemNetAlias(c *CoreItemNetAliasContext)

	// EnterCoreItemInitial is called when entering the CoreItemInitial production.
	EnterCoreItemInitial(c *CoreItemInitialContext)

	// EnterCoreItemFinal is called when entering the CoreItemFinal production.
	EnterCoreItemFinal(c *CoreItemFinalContext)

	// EnterCoreItemAlways is called when entering the CoreItemAlways production.
	EnterCoreItemAlways(c *CoreItemAlwaysContext)

	// EnterCoreItemLoop is called when entering the CoreItemLoop production.
	EnterCoreItemLoop(c *CoreItemLoopContext)

	// EnterCoreItemCondGen is called when entering the CoreItemCondGen production.
	EnterCoreItemCondGen(c *CoreItemCondGenContext)

	// EnterCoreItemElab is called when entering the CoreItemElab production.
	EnterCoreItemElab(c *CoreItemElabContext)

	// EnterModule_item_item is called when entering the module_item_item production.
	EnterModule_item_item(c *Module_item_itemContext)

	// EnterModuleGenReg is called when entering the ModuleGenReg production.
	EnterModuleGenReg(c *ModuleGenRegContext)

	// EnterModuleItemItem is called when entering the ModuleItemItem production.
	EnterModuleItemItem(c *ModuleItemItemContext)

	// EnterModuleSpecBlock is called when entering the ModuleSpecBlock production.
	EnterModuleSpecBlock(c *ModuleSpecBlockContext)

	// EnterModuleProgDecl is called when entering the ModuleProgDecl production.
	EnterModuleProgDecl(c *ModuleProgDeclContext)

	// EnterModuleDecl is called when entering the ModuleDecl production.
	EnterModuleDecl(c *ModuleDeclContext)

	// EnterModuleIntfDecl is called when entering the ModuleIntfDecl production.
	EnterModuleIntfDecl(c *ModuleIntfDeclContext)

	// EnterModuleTimeDecl is called when entering the ModuleTimeDecl production.
	EnterModuleTimeDecl(c *ModuleTimeDeclContext)

	// EnterModulePortDecl is called when entering the ModulePortDecl production.
	EnterModulePortDecl(c *ModulePortDeclContext)

	// ExitSource_text is called when exiting the source_text production.
	ExitSource_text(c *Source_textContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitHeader_text is called when exiting the header_text production.
	ExitHeader_text(c *Header_textContext)

	// ExitDesign_attribute is called when exiting the design_attribute production.
	ExitDesign_attribute(c *Design_attributeContext)

	// ExitCompiler_directive is called when exiting the compiler_directive production.
	ExitCompiler_directive(c *Compiler_directiveContext)

	// ExitTime_num is called when exiting the time_num production.
	ExitTime_num(c *Time_numContext)

	// ExitTime_lit is called when exiting the time_lit production.
	ExitTime_lit(c *Time_litContext)

	// ExitTimescale_compiler_directive is called when exiting the timescale_compiler_directive production.
	ExitTimescale_compiler_directive(c *Timescale_compiler_directiveContext)

	// ExitDefault_nettype_statement is called when exiting the default_nettype_statement production.
	ExitDefault_nettype_statement(c *Default_nettype_statementContext)

	// ExitInclude_svh is called when exiting the include_svh production.
	ExitInclude_svh(c *Include_svhContext)

	// ExitAssignment_operator is called when exiting the assignment_operator production.
	ExitAssignment_operator(c *Assignment_operatorContext)

	// ExitEdge_identifier is called when exiting the edge_identifier production.
	ExitEdge_identifier(c *Edge_identifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitInteger_type is called when exiting the integer_type production.
	ExitInteger_type(c *Integer_typeContext)

	// ExitInteger_atom_type is called when exiting the integer_atom_type production.
	ExitInteger_atom_type(c *Integer_atom_typeContext)

	// ExitInteger_vector_type is called when exiting the integer_vector_type production.
	ExitInteger_vector_type(c *Integer_vector_typeContext)

	// ExitNon_integer_type is called when exiting the non_integer_type production.
	ExitNon_integer_type(c *Non_integer_typeContext)

	// ExitNet_type is called when exiting the net_type production.
	ExitNet_type(c *Net_typeContext)

	// ExitUnary_module_path_operator is called when exiting the unary_module_path_operator production.
	ExitUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitInc_or_dec_operator is called when exiting the inc_or_dec_operator production.
	ExitInc_or_dec_operator(c *Inc_or_dec_operatorContext)

	// ExitImplicit_class_handle is called when exiting the implicit_class_handle production.
	ExitImplicit_class_handle(c *Implicit_class_handleContext)

	// ExitIntegral_number is called when exiting the integral_number production.
	ExitIntegral_number(c *Integral_numberContext)

	// ExitReal_number is called when exiting the real_number production.
	ExitReal_number(c *Real_numberContext)

	// ExitAny_system_tf_identifier is called when exiting the any_system_tf_identifier production.
	ExitAny_system_tf_identifier(c *Any_system_tf_identifierContext)

	// ExitSigning is called when exiting the signing production.
	ExitSigning(c *SigningContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitTimeunits_declaration is called when exiting the timeunits_declaration production.
	ExitTimeunits_declaration(c *Timeunits_declarationContext)

	// ExitLifetime is called when exiting the lifetime production.
	ExitLifetime(c *LifetimeContext)

	// ExitPort_direction is called when exiting the port_direction production.
	ExitPort_direction(c *Port_directionContext)

	// ExitAlways_keyword is called when exiting the always_keyword production.
	ExitAlways_keyword(c *Always_keywordContext)

	// ExitJoin_keyword is called when exiting the join_keyword production.
	ExitJoin_keyword(c *Join_keywordContext)

	// ExitUnique_priority is called when exiting the unique_priority production.
	ExitUnique_priority(c *Unique_priorityContext)

	// ExitDrive_strength is called when exiting the drive_strength production.
	ExitDrive_strength(c *Drive_strengthContext)

	// ExitStrength0 is called when exiting the strength0 production.
	ExitStrength0(c *Strength0Context)

	// ExitStrength1 is called when exiting the strength1 production.
	ExitStrength1(c *Strength1Context)

	// ExitCharge_strength is called when exiting the charge_strength production.
	ExitCharge_strength(c *Charge_strengthContext)

	// ExitSequence_lvar_port_direction is called when exiting the sequence_lvar_port_direction production.
	ExitSequence_lvar_port_direction(c *Sequence_lvar_port_directionContext)

	// ExitBins_keyword is called when exiting the bins_keyword production.
	ExitBins_keyword(c *Bins_keywordContext)

	// ExitClass_item_qualifier is called when exiting the class_item_qualifier production.
	ExitClass_item_qualifier(c *Class_item_qualifierContext)

	// ExitRandom_qualifier is called when exiting the random_qualifier production.
	ExitRandom_qualifier(c *Random_qualifierContext)

	// ExitProperty_qualifier is called when exiting the property_qualifier production.
	ExitProperty_qualifier(c *Property_qualifierContext)

	// ExitMethod_qualifier is called when exiting the method_qualifier production.
	ExitMethod_qualifier(c *Method_qualifierContext)

	// ExitConstraint_prototype_qualifier is called when exiting the constraint_prototype_qualifier production.
	ExitConstraint_prototype_qualifier(c *Constraint_prototype_qualifierContext)

	// ExitCmos_switchtype is called when exiting the cmos_switchtype production.
	ExitCmos_switchtype(c *Cmos_switchtypeContext)

	// ExitEnable_gatetype is called when exiting the enable_gatetype production.
	ExitEnable_gatetype(c *Enable_gatetypeContext)

	// ExitMos_switchtype is called when exiting the mos_switchtype production.
	ExitMos_switchtype(c *Mos_switchtypeContext)

	// ExitN_input_gatetype is called when exiting the n_input_gatetype production.
	ExitN_input_gatetype(c *N_input_gatetypeContext)

	// ExitN_output_gatetype is called when exiting the n_output_gatetype production.
	ExitN_output_gatetype(c *N_output_gatetypeContext)

	// ExitPass_en_switchtype is called when exiting the pass_en_switchtype production.
	ExitPass_en_switchtype(c *Pass_en_switchtypeContext)

	// ExitPass_switchtype is called when exiting the pass_switchtype production.
	ExitPass_switchtype(c *Pass_switchtypeContext)

	// ExitAny_implication is called when exiting the any_implication production.
	ExitAny_implication(c *Any_implicationContext)

	// ExitPolarity_operator is called when exiting the polarity_operator production.
	ExitPolarity_operator(c *Polarity_operatorContext)

	// ExitTiming_check_event_control is called when exiting the timing_check_event_control production.
	ExitTiming_check_event_control(c *Timing_check_event_controlContext)

	// ExitImport_export is called when exiting the import_export production.
	ExitImport_export(c *Import_exportContext)

	// ExitArray_method_name is called when exiting the array_method_name production.
	ExitArray_method_name(c *Array_method_nameContext)

	// ExitOperator_mul_div_mod is called when exiting the operator_mul_div_mod production.
	ExitOperator_mul_div_mod(c *Operator_mul_div_modContext)

	// ExitOperator_plus_minus is called when exiting the operator_plus_minus production.
	ExitOperator_plus_minus(c *Operator_plus_minusContext)

	// ExitOperator_shift is called when exiting the operator_shift production.
	ExitOperator_shift(c *Operator_shiftContext)

	// ExitOperator_cmp is called when exiting the operator_cmp production.
	ExitOperator_cmp(c *Operator_cmpContext)

	// ExitOperator_eq_neq is called when exiting the operator_eq_neq production.
	ExitOperator_eq_neq(c *Operator_eq_neqContext)

	// ExitOperator_xor is called when exiting the operator_xor production.
	ExitOperator_xor(c *Operator_xorContext)

	// ExitOperator_impl is called when exiting the operator_impl production.
	ExitOperator_impl(c *Operator_implContext)

	// ExitUdp_nonansi_declaration is called when exiting the udp_nonansi_declaration production.
	ExitUdp_nonansi_declaration(c *Udp_nonansi_declarationContext)

	// ExitUdp_ansi_declaration is called when exiting the udp_ansi_declaration production.
	ExitUdp_ansi_declaration(c *Udp_ansi_declarationContext)

	// ExitUdp_declaration is called when exiting the udp_declaration production.
	ExitUdp_declaration(c *Udp_declarationContext)

	// ExitUdp_declaration_port_list is called when exiting the udp_declaration_port_list production.
	ExitUdp_declaration_port_list(c *Udp_declaration_port_listContext)

	// ExitUdp_port_declaration is called when exiting the udp_port_declaration production.
	ExitUdp_port_declaration(c *Udp_port_declarationContext)

	// ExitUdp_output_declaration is called when exiting the udp_output_declaration production.
	ExitUdp_output_declaration(c *Udp_output_declarationContext)

	// ExitUdp_input_declaration is called when exiting the udp_input_declaration production.
	ExitUdp_input_declaration(c *Udp_input_declarationContext)

	// ExitUdp_reg_declaration is called when exiting the udp_reg_declaration production.
	ExitUdp_reg_declaration(c *Udp_reg_declarationContext)

	// ExitUdp_body is called when exiting the udp_body production.
	ExitUdp_body(c *Udp_bodyContext)

	// ExitCombinational_body is called when exiting the combinational_body production.
	ExitCombinational_body(c *Combinational_bodyContext)

	// ExitCombinational_entry is called when exiting the combinational_entry production.
	ExitCombinational_entry(c *Combinational_entryContext)

	// ExitSequential_body is called when exiting the sequential_body production.
	ExitSequential_body(c *Sequential_bodyContext)

	// ExitUdp_initial_statement is called when exiting the udp_initial_statement production.
	ExitUdp_initial_statement(c *Udp_initial_statementContext)

	// ExitSequential_entry is called when exiting the sequential_entry production.
	ExitSequential_entry(c *Sequential_entryContext)

	// ExitSeq_input_list is called when exiting the seq_input_list production.
	ExitSeq_input_list(c *Seq_input_listContext)

	// ExitLevel_input_list is called when exiting the level_input_list production.
	ExitLevel_input_list(c *Level_input_listContext)

	// ExitEdge_input_list is called when exiting the edge_input_list production.
	ExitEdge_input_list(c *Edge_input_listContext)

	// ExitEdge_indicator is called when exiting the edge_indicator production.
	ExitEdge_indicator(c *Edge_indicatorContext)

	// ExitCurrent_state is called when exiting the current_state production.
	ExitCurrent_state(c *Current_stateContext)

	// ExitNext_state is called when exiting the next_state production.
	ExitNext_state(c *Next_stateContext)

	// ExitInterface_declaration is called when exiting the interface_declaration production.
	ExitInterface_declaration(c *Interface_declarationContext)

	// ExitInterface_header is called when exiting the interface_header production.
	ExitInterface_header(c *Interface_headerContext)

	// ExitInterface_item is called when exiting the interface_item production.
	ExitInterface_item(c *Interface_itemContext)

	// ExitModport_declaration is called when exiting the modport_declaration production.
	ExitModport_declaration(c *Modport_declarationContext)

	// ExitModport_item is called when exiting the modport_item production.
	ExitModport_item(c *Modport_itemContext)

	// ExitModport_ports_declaration is called when exiting the modport_ports_declaration production.
	ExitModport_ports_declaration(c *Modport_ports_declarationContext)

	// ExitModport_clocking_declaration is called when exiting the modport_clocking_declaration production.
	ExitModport_clocking_declaration(c *Modport_clocking_declarationContext)

	// ExitModport_simple_ports_declaration is called when exiting the modport_simple_ports_declaration production.
	ExitModport_simple_ports_declaration(c *Modport_simple_ports_declarationContext)

	// ExitModport_simple_port is called when exiting the modport_simple_port production.
	ExitModport_simple_port(c *Modport_simple_portContext)

	// ExitModport_tf_ports_declaration is called when exiting the modport_tf_ports_declaration production.
	ExitModport_tf_ports_declaration(c *Modport_tf_ports_declarationContext)

	// ExitModport_tf_port is called when exiting the modport_tf_port production.
	ExitModport_tf_port(c *Modport_tf_portContext)

	// ExitStatement_or_null is called when exiting the statement_or_null production.
	ExitStatement_or_null(c *Statement_or_nullContext)

	// ExitInitial_construct is called when exiting the initial_construct production.
	ExitInitial_construct(c *Initial_constructContext)

	// ExitDefault_clocking_or_dissable_construct is called when exiting the default_clocking_or_dissable_construct production.
	ExitDefault_clocking_or_dissable_construct(c *Default_clocking_or_dissable_constructContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStmtItemMain is called when exiting the StmtItemMain production.
	ExitStmtItemMain(c *StmtItemMainContext)

	// ExitStmtItemMacro is called when exiting the StmtItemMacro production.
	ExitStmtItemMacro(c *StmtItemMacroContext)

	// ExitStmtItemCase is called when exiting the StmtItemCase production.
	ExitStmtItemCase(c *StmtItemCaseContext)

	// ExitStmtItemCond is called when exiting the StmtItemCond production.
	ExitStmtItemCond(c *StmtItemCondContext)

	// ExitStmtItemSubCall is called when exiting the StmtItemSubCall production.
	ExitStmtItemSubCall(c *StmtItemSubCallContext)

	// ExitStmtItemDisable is called when exiting the StmtItemDisable production.
	ExitStmtItemDisable(c *StmtItemDisableContext)

	// ExitStmtItemEvent is called when exiting the StmtItemEvent production.
	ExitStmtItemEvent(c *StmtItemEventContext)

	// ExitStmtItemLoop is called when exiting the StmtItemLoop production.
	ExitStmtItemLoop(c *StmtItemLoopContext)

	// ExitStmtItemJump is called when exiting the StmtItemJump production.
	ExitStmtItemJump(c *StmtItemJumpContext)

	// ExitStmtItemPar is called when exiting the StmtItemPar production.
	ExitStmtItemPar(c *StmtItemParContext)

	// ExitStmtItemProcTime is called when exiting the StmtItemProcTime production.
	ExitStmtItemProcTime(c *StmtItemProcTimeContext)

	// ExitStmtItemSeq is called when exiting the StmtItemSeq production.
	ExitStmtItemSeq(c *StmtItemSeqContext)

	// ExitStmtItemWait is called when exiting the StmtItemWait production.
	ExitStmtItemWait(c *StmtItemWaitContext)

	// ExitStmtItemProcAssert is called when exiting the StmtItemProcAssert production.
	ExitStmtItemProcAssert(c *StmtItemProcAssertContext)

	// ExitStmtItemRandseq is called when exiting the StmtItemRandseq production.
	ExitStmtItemRandseq(c *StmtItemRandseqContext)

	// ExitStmtItemRandcase is called when exiting the StmtItemRandcase production.
	ExitStmtItemRandcase(c *StmtItemRandcaseContext)

	// ExitStmtItemExpect is called when exiting the StmtItemExpect production.
	ExitStmtItemExpect(c *StmtItemExpectContext)

	// ExitMacro_statement is called when exiting the macro_statement production.
	ExitMacro_statement(c *Macro_statementContext)

	// ExitCycle_delay is called when exiting the cycle_delay production.
	ExitCycle_delay(c *Cycle_delayContext)

	// ExitClocking_drive is called when exiting the clocking_drive production.
	ExitClocking_drive(c *Clocking_driveContext)

	// ExitClockvar_expression is called when exiting the clockvar_expression production.
	ExitClockvar_expression(c *Clockvar_expressionContext)

	// ExitFinal_construct is called when exiting the final_construct production.
	ExitFinal_construct(c *Final_constructContext)

	// ExitBlocking_assignment is called when exiting the blocking_assignment production.
	ExitBlocking_assignment(c *Blocking_assignmentContext)

	// ExitProcedural_timing_control_statement is called when exiting the procedural_timing_control_statement production.
	ExitProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// ExitProcedural_timing_control is called when exiting the procedural_timing_control production.
	ExitProcedural_timing_control(c *Procedural_timing_controlContext)

	// ExitEvent_control is called when exiting the event_control production.
	ExitEvent_control(c *Event_controlContext)

	// ExitDelay_or_event_control is called when exiting the delay_or_event_control production.
	ExitDelay_or_event_control(c *Delay_or_event_controlContext)

	// ExitDelay3 is called when exiting the delay3 production.
	ExitDelay3(c *Delay3Context)

	// ExitDelay2 is called when exiting the delay2 production.
	ExitDelay2(c *Delay2Context)

	// ExitDelay_value is called when exiting the delay_value production.
	ExitDelay_value(c *Delay_valueContext)

	// ExitDelay_control is called when exiting the delay_control production.
	ExitDelay_control(c *Delay_controlContext)

	// ExitNonblocking_assignment is called when exiting the nonblocking_assignment production.
	ExitNonblocking_assignment(c *Nonblocking_assignmentContext)

	// ExitProcedural_continuous_assignment is called when exiting the procedural_continuous_assignment production.
	ExitProcedural_continuous_assignment(c *Procedural_continuous_assignmentContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitAction_block is called when exiting the action_block production.
	ExitAction_block(c *Action_blockContext)

	// ExitSeq_block is called when exiting the seq_block production.
	ExitSeq_block(c *Seq_blockContext)

	// ExitPar_block is called when exiting the par_block production.
	ExitPar_block(c *Par_blockContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_keyword is called when exiting the case_keyword production.
	ExitCase_keyword(c *Case_keywordContext)

	// ExitCase_item is called when exiting the case_item production.
	ExitCase_item(c *Case_itemContext)

	// ExitCase_pattern_item is called when exiting the case_pattern_item production.
	ExitCase_pattern_item(c *Case_pattern_itemContext)

	// ExitCase_inside_item is called when exiting the case_inside_item production.
	ExitCase_inside_item(c *Case_inside_itemContext)

	// ExitRandcase_statement is called when exiting the randcase_statement production.
	ExitRandcase_statement(c *Randcase_statementContext)

	// ExitRandcase_item is called when exiting the randcase_item production.
	ExitRandcase_item(c *Randcase_itemContext)

	// ExitCond_predicate is called when exiting the cond_predicate production.
	ExitCond_predicate(c *Cond_predicateContext)

	// ExitConditional_statement is called when exiting the conditional_statement production.
	ExitConditional_statement(c *Conditional_statementContext)

	// ExitSubroutine_call_statement is called when exiting the subroutine_call_statement production.
	ExitSubroutine_call_statement(c *Subroutine_call_statementContext)

	// ExitDisable_statement is called when exiting the disable_statement production.
	ExitDisable_statement(c *Disable_statementContext)

	// ExitEvent_trigger is called when exiting the event_trigger production.
	ExitEvent_trigger(c *Event_triggerContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitList_of_variable_assignments is called when exiting the list_of_variable_assignments production.
	ExitList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// ExitFor_initialization is called when exiting the for_initialization production.
	ExitFor_initialization(c *For_initializationContext)

	// ExitFor_variable_declaration_var_assign is called when exiting the for_variable_declaration_var_assign production.
	ExitFor_variable_declaration_var_assign(c *For_variable_declaration_var_assignContext)

	// ExitFor_variable_declaration is called when exiting the for_variable_declaration production.
	ExitFor_variable_declaration(c *For_variable_declarationContext)

	// ExitFor_step is called when exiting the for_step production.
	ExitFor_step(c *For_stepContext)

	// ExitLoop_variables is called when exiting the loop_variables production.
	ExitLoop_variables(c *Loop_variablesContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitName_of_instance is called when exiting the name_of_instance production.
	ExitName_of_instance(c *Name_of_instanceContext)

	// ExitChecker_instantiation is called when exiting the checker_instantiation production.
	ExitChecker_instantiation(c *Checker_instantiationContext)

	// ExitList_of_checker_port_connections is called when exiting the list_of_checker_port_connections production.
	ExitList_of_checker_port_connections(c *List_of_checker_port_connectionsContext)

	// ExitOrdered_checker_port_connection is called when exiting the ordered_checker_port_connection production.
	ExitOrdered_checker_port_connection(c *Ordered_checker_port_connectionContext)

	// ExitNamed_checker_port_connection is called when exiting the named_checker_port_connection production.
	ExitNamed_checker_port_connection(c *Named_checker_port_connectionContext)

	// ExitProcedural_assertion_statement is called when exiting the procedural_assertion_statement production.
	ExitProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// ExitConcurrent_assertion_statement is called when exiting the concurrent_assertion_statement production.
	ExitConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// ExitAssertion_item is called when exiting the assertion_item production.
	ExitAssertion_item(c *Assertion_itemContext)

	// ExitConcurrent_assertion_item is called when exiting the concurrent_assertion_item production.
	ExitConcurrent_assertion_item(c *Concurrent_assertion_itemContext)

	// ExitImmediate_assertion_statement is called when exiting the immediate_assertion_statement production.
	ExitImmediate_assertion_statement(c *Immediate_assertion_statementContext)

	// ExitSimple_immediate_assertion_statement is called when exiting the simple_immediate_assertion_statement production.
	ExitSimple_immediate_assertion_statement(c *Simple_immediate_assertion_statementContext)

	// ExitSimple_immediate_assert_statement is called when exiting the simple_immediate_assert_statement production.
	ExitSimple_immediate_assert_statement(c *Simple_immediate_assert_statementContext)

	// ExitSimple_immediate_assume_statement is called when exiting the simple_immediate_assume_statement production.
	ExitSimple_immediate_assume_statement(c *Simple_immediate_assume_statementContext)

	// ExitSimple_immediate_cover_statement is called when exiting the simple_immediate_cover_statement production.
	ExitSimple_immediate_cover_statement(c *Simple_immediate_cover_statementContext)

	// ExitDeferred_immediate_assertion_statement is called when exiting the deferred_immediate_assertion_statement production.
	ExitDeferred_immediate_assertion_statement(c *Deferred_immediate_assertion_statementContext)

	// ExitPrimitive_delay is called when exiting the primitive_delay production.
	ExitPrimitive_delay(c *Primitive_delayContext)

	// ExitDeferred_immediate_assert_statement is called when exiting the deferred_immediate_assert_statement production.
	ExitDeferred_immediate_assert_statement(c *Deferred_immediate_assert_statementContext)

	// ExitDeferred_immediate_assume_statement is called when exiting the deferred_immediate_assume_statement production.
	ExitDeferred_immediate_assume_statement(c *Deferred_immediate_assume_statementContext)

	// ExitDeferred_immediate_cover_statement is called when exiting the deferred_immediate_cover_statement production.
	ExitDeferred_immediate_cover_statement(c *Deferred_immediate_cover_statementContext)

	// ExitWeight_specification is called when exiting the weight_specification production.
	ExitWeight_specification(c *Weight_specificationContext)

	// ExitProduction_item is called when exiting the production_item production.
	ExitProduction_item(c *Production_itemContext)

	// ExitRs_code_block is called when exiting the rs_code_block production.
	ExitRs_code_block(c *Rs_code_blockContext)

	// ExitRandsequence_statement is called when exiting the randsequence_statement production.
	ExitRandsequence_statement(c *Randsequence_statementContext)

	// ExitRs_prod is called when exiting the rs_prod production.
	ExitRs_prod(c *Rs_prodContext)

	// ExitRs_if_else is called when exiting the rs_if_else production.
	ExitRs_if_else(c *Rs_if_elseContext)

	// ExitRs_repeat is called when exiting the rs_repeat production.
	ExitRs_repeat(c *Rs_repeatContext)

	// ExitRs_case is called when exiting the rs_case production.
	ExitRs_case(c *Rs_caseContext)

	// ExitRs_case_item is called when exiting the rs_case_item production.
	ExitRs_case_item(c *Rs_case_itemContext)

	// ExitRs_rule is called when exiting the rs_rule production.
	ExitRs_rule(c *Rs_ruleContext)

	// ExitRs_production_list is called when exiting the rs_production_list production.
	ExitRs_production_list(c *Rs_production_listContext)

	// ExitProduction is called when exiting the production production.
	ExitProduction(c *ProductionContext)

	// ExitTf_item_declaration is called when exiting the tf_item_declaration production.
	ExitTf_item_declaration(c *Tf_item_declarationContext)

	// ExitTf_port_list is called when exiting the tf_port_list production.
	ExitTf_port_list(c *Tf_port_listContext)

	// ExitTf_port_item is called when exiting the tf_port_item production.
	ExitTf_port_item(c *Tf_port_itemContext)

	// ExitTf_port_direction is called when exiting the tf_port_direction production.
	ExitTf_port_direction(c *Tf_port_directionContext)

	// ExitTf_port_declaration is called when exiting the tf_port_declaration production.
	ExitTf_port_declaration(c *Tf_port_declarationContext)

	// ExitList_of_tf_variable_identifiers_item is called when exiting the list_of_tf_variable_identifiers_item production.
	ExitList_of_tf_variable_identifiers_item(c *List_of_tf_variable_identifiers_itemContext)

	// ExitList_of_tf_variable_identifiers is called when exiting the list_of_tf_variable_identifiers production.
	ExitList_of_tf_variable_identifiers(c *List_of_tf_variable_identifiersContext)

	// ExitExpect_property_statement is called when exiting the expect_property_statement production.
	ExitExpect_property_statement(c *Expect_property_statementContext)

	// ExitBlock_item_declaration is called when exiting the block_item_declaration production.
	ExitBlock_item_declaration(c *Block_item_declarationContext)

	// ExitParam_assignment is called when exiting the param_assignment production.
	ExitParam_assignment(c *Param_assignmentContext)

	// ExitType_assignment is called when exiting the type_assignment production.
	ExitType_assignment(c *Type_assignmentContext)

	// ExitList_of_type_assignments is called when exiting the list_of_type_assignments production.
	ExitList_of_type_assignments(c *List_of_type_assignmentsContext)

	// ExitList_of_param_assignments is called when exiting the list_of_param_assignments production.
	ExitList_of_param_assignments(c *List_of_param_assignmentsContext)

	// ExitParameter_declaration_primitive is called when exiting the parameter_declaration_primitive production.
	ExitParameter_declaration_primitive(c *Parameter_declaration_primitiveContext)

	// ExitLocal_parameter_declaration is called when exiting the local_parameter_declaration production.
	ExitLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitNet_type_declaration is called when exiting the net_type_declaration production.
	ExitNet_type_declaration(c *Net_type_declarationContext)

	// ExitLet_declaration is called when exiting the let_declaration production.
	ExitLet_declaration(c *Let_declarationContext)

	// ExitLet_port_list is called when exiting the let_port_list production.
	ExitLet_port_list(c *Let_port_listContext)

	// ExitLet_port_item is called when exiting the let_port_item production.
	ExitLet_port_item(c *Let_port_itemContext)

	// ExitLet_formal_type is called when exiting the let_formal_type production.
	ExitLet_formal_type(c *Let_formal_typeContext)

	// ExitPackage_import_declaration is called when exiting the package_import_declaration production.
	ExitPackage_import_declaration(c *Package_import_declarationContext)

	// ExitPackage_import_item is called when exiting the package_import_item production.
	ExitPackage_import_item(c *Package_import_itemContext)

	// ExitProperty_list_of_arguments is called when exiting the property_list_of_arguments production.
	ExitProperty_list_of_arguments(c *Property_list_of_argumentsContext)

	// ExitProperty_actual_arg is called when exiting the property_actual_arg production.
	ExitProperty_actual_arg(c *Property_actual_argContext)

	// ExitProperty_formal_type is called when exiting the property_formal_type production.
	ExitProperty_formal_type(c *Property_formal_typeContext)

	// ExitSequence_formal_type is called when exiting the sequence_formal_type production.
	ExitSequence_formal_type(c *Sequence_formal_typeContext)

	// ExitProperty_instance is called when exiting the property_instance production.
	ExitProperty_instance(c *Property_instanceContext)

	// ExitProperty_spec is called when exiting the property_spec production.
	ExitProperty_spec(c *Property_specContext)

	// ExitProperty_expr is called when exiting the property_expr production.
	ExitProperty_expr(c *Property_exprContext)

	// ExitProperty_case_item is called when exiting the property_case_item production.
	ExitProperty_case_item(c *Property_case_itemContext)

	// ExitBit_select is called when exiting the bit_select production.
	ExitBit_select(c *Bit_selectContext)

	// ExitIdentifier_with_bit_select is called when exiting the identifier_with_bit_select production.
	ExitIdentifier_with_bit_select(c *Identifier_with_bit_selectContext)

	// ExitPackage_or_class_scoped_hier_id_with_select is called when exiting the package_or_class_scoped_hier_id_with_select production.
	ExitPackage_or_class_scoped_hier_id_with_select(c *Package_or_class_scoped_hier_id_with_selectContext)

	// ExitPackage_or_class_scoped_path_item is called when exiting the package_or_class_scoped_path_item production.
	ExitPackage_or_class_scoped_path_item(c *Package_or_class_scoped_path_itemContext)

	// ExitPackage_or_class_scoped_path is called when exiting the package_or_class_scoped_path production.
	ExitPackage_or_class_scoped_path(c *Package_or_class_scoped_pathContext)

	// ExitHierarchical_identifier is called when exiting the hierarchical_identifier production.
	ExitHierarchical_identifier(c *Hierarchical_identifierContext)

	// ExitPackage_or_class_scoped_id is called when exiting the package_or_class_scoped_id production.
	ExitPackage_or_class_scoped_id(c *Package_or_class_scoped_idContext)

	// ExitSelect_ is called when exiting the select_ production.
	ExitSelect_(c *Select_Context)

	// ExitEvent_expression_item is called when exiting the event_expression_item production.
	ExitEvent_expression_item(c *Event_expression_itemContext)

	// ExitEvent_expression is called when exiting the event_expression production.
	ExitEvent_expression(c *Event_expressionContext)

	// ExitBoolean_abbrev is called when exiting the boolean_abbrev production.
	ExitBoolean_abbrev(c *Boolean_abbrevContext)

	// ExitSequence_abbrev is called when exiting the sequence_abbrev production.
	ExitSequence_abbrev(c *Sequence_abbrevContext)

	// ExitConsecutive_repetition is called when exiting the consecutive_repetition production.
	ExitConsecutive_repetition(c *Consecutive_repetitionContext)

	// ExitNon_consecutive_repetition is called when exiting the non_consecutive_repetition production.
	ExitNon_consecutive_repetition(c *Non_consecutive_repetitionContext)

	// ExitGoto_repetition is called when exiting the goto_repetition production.
	ExitGoto_repetition(c *Goto_repetitionContext)

	// ExitCycle_delay_const_range_expression is called when exiting the cycle_delay_const_range_expression production.
	ExitCycle_delay_const_range_expression(c *Cycle_delay_const_range_expressionContext)

	// ExitSequence_instance is called when exiting the sequence_instance production.
	ExitSequence_instance(c *Sequence_instanceContext)

	// ExitSequence_expr is called when exiting the sequence_expr production.
	ExitSequence_expr(c *Sequence_exprContext)

	// ExitSequence_match_item is called when exiting the sequence_match_item production.
	ExitSequence_match_item(c *Sequence_match_itemContext)

	// ExitOperator_assignment is called when exiting the operator_assignment production.
	ExitOperator_assignment(c *Operator_assignmentContext)

	// ExitSequence_actual_arg is called when exiting the sequence_actual_arg production.
	ExitSequence_actual_arg(c *Sequence_actual_argContext)

	// ExitDist_weight is called when exiting the dist_weight production.
	ExitDist_weight(c *Dist_weightContext)

	// ExitClocking_declaration is called when exiting the clocking_declaration production.
	ExitClocking_declaration(c *Clocking_declarationContext)

	// ExitClocking_item is called when exiting the clocking_item production.
	ExitClocking_item(c *Clocking_itemContext)

	// ExitList_of_clocking_decl_assign is called when exiting the list_of_clocking_decl_assign production.
	ExitList_of_clocking_decl_assign(c *List_of_clocking_decl_assignContext)

	// ExitClocking_decl_assign is called when exiting the clocking_decl_assign production.
	ExitClocking_decl_assign(c *Clocking_decl_assignContext)

	// ExitDefault_skew is called when exiting the default_skew production.
	ExitDefault_skew(c *Default_skewContext)

	// ExitClocking_direction is called when exiting the clocking_direction production.
	ExitClocking_direction(c *Clocking_directionContext)

	// ExitClocking_skew is called when exiting the clocking_skew production.
	ExitClocking_skew(c *Clocking_skewContext)

	// ExitClocking_event is called when exiting the clocking_event production.
	ExitClocking_event(c *Clocking_eventContext)

	// ExitCycle_delay_range is called when exiting the cycle_delay_range production.
	ExitCycle_delay_range(c *Cycle_delay_rangeContext)

	// ExitExpression_or_dist is called when exiting the expression_or_dist production.
	ExitExpression_or_dist(c *Expression_or_distContext)

	// ExitCovergroup_declaration is called when exiting the covergroup_declaration production.
	ExitCovergroup_declaration(c *Covergroup_declarationContext)

	// ExitCover_cross is called when exiting the cover_cross production.
	ExitCover_cross(c *Cover_crossContext)

	// ExitIdentifier_list_2plus is called when exiting the identifier_list_2plus production.
	ExitIdentifier_list_2plus(c *Identifier_list_2plusContext)

	// ExitCross_body is called when exiting the cross_body production.
	ExitCross_body(c *Cross_bodyContext)

	// ExitCross_body_item is called when exiting the cross_body_item production.
	ExitCross_body_item(c *Cross_body_itemContext)

	// ExitBins_selection_or_option is called when exiting the bins_selection_or_option production.
	ExitBins_selection_or_option(c *Bins_selection_or_optionContext)

	// ExitBins_selection is called when exiting the bins_selection production.
	ExitBins_selection(c *Bins_selectionContext)

	// ExitSelect_expression is called when exiting the select_expression production.
	ExitSelect_expression(c *Select_expressionContext)

	// ExitSelect_condition is called when exiting the select_condition production.
	ExitSelect_condition(c *Select_conditionContext)

	// ExitBins_expression is called when exiting the bins_expression production.
	ExitBins_expression(c *Bins_expressionContext)

	// ExitCovergroup_range_list is called when exiting the covergroup_range_list production.
	ExitCovergroup_range_list(c *Covergroup_range_listContext)

	// ExitCovergroup_value_range is called when exiting the covergroup_value_range production.
	ExitCovergroup_value_range(c *Covergroup_value_rangeContext)

	// ExitCovergroup_expression is called when exiting the covergroup_expression production.
	ExitCovergroup_expression(c *Covergroup_expressionContext)

	// ExitCoverage_spec_or_option is called when exiting the coverage_spec_or_option production.
	ExitCoverage_spec_or_option(c *Coverage_spec_or_optionContext)

	// ExitCoverage_option is called when exiting the coverage_option production.
	ExitCoverage_option(c *Coverage_optionContext)

	// ExitCoverage_spec is called when exiting the coverage_spec production.
	ExitCoverage_spec(c *Coverage_specContext)

	// ExitCover_point is called when exiting the cover_point production.
	ExitCover_point(c *Cover_pointContext)

	// ExitBins_or_empty is called when exiting the bins_or_empty production.
	ExitBins_or_empty(c *Bins_or_emptyContext)

	// ExitBins_or_options is called when exiting the bins_or_options production.
	ExitBins_or_options(c *Bins_or_optionsContext)

	// ExitTrans_list is called when exiting the trans_list production.
	ExitTrans_list(c *Trans_listContext)

	// ExitTrans_set is called when exiting the trans_set production.
	ExitTrans_set(c *Trans_setContext)

	// ExitTrans_range_list is called when exiting the trans_range_list production.
	ExitTrans_range_list(c *Trans_range_listContext)

	// ExitRepeat_range is called when exiting the repeat_range production.
	ExitRepeat_range(c *Repeat_rangeContext)

	// ExitCoverage_event is called when exiting the coverage_event production.
	ExitCoverage_event(c *Coverage_eventContext)

	// ExitBlock_event_expression is called when exiting the block_event_expression production.
	ExitBlock_event_expression(c *Block_event_expressionContext)

	// ExitHierarchical_btf_identifier is called when exiting the hierarchical_btf_identifier production.
	ExitHierarchical_btf_identifier(c *Hierarchical_btf_identifierContext)

	// ExitAssertion_variable_declaration is called when exiting the assertion_variable_declaration production.
	ExitAssertion_variable_declaration(c *Assertion_variable_declarationContext)

	// ExitDist_item is called when exiting the dist_item production.
	ExitDist_item(c *Dist_itemContext)

	// ExitValue_range is called when exiting the value_range production.
	ExitValue_range(c *Value_rangeContext)

	// ExitAttribute_instance is called when exiting the attribute_instance production.
	ExitAttribute_instance(c *Attribute_instanceContext)

	// ExitAttr_spec is called when exiting the attr_spec production.
	ExitAttr_spec(c *Attr_specContext)

	// ExitClass_new is called when exiting the class_new production.
	ExitClass_new(c *Class_newContext)

	// ExitParam_expression is called when exiting the param_expression production.
	ExitParam_expression(c *Param_expressionContext)

	// ExitConstant_param_expression is called when exiting the constant_param_expression production.
	ExitConstant_param_expression(c *Constant_param_expressionContext)

	// ExitUnpacked_dimension is called when exiting the unpacked_dimension production.
	ExitUnpacked_dimension(c *Unpacked_dimensionContext)

	// ExitPacked_dimension is called when exiting the packed_dimension production.
	ExitPacked_dimension(c *Packed_dimensionContext)

	// ExitVariable_dimension is called when exiting the variable_dimension production.
	ExitVariable_dimension(c *Variable_dimensionContext)

	// ExitStruct_union is called when exiting the struct_union production.
	ExitStruct_union(c *Struct_unionContext)

	// ExitEnum_base_type is called when exiting the enum_base_type production.
	ExitEnum_base_type(c *Enum_base_typeContext)

	// ExitData_type_primitive is called when exiting the data_type_primitive production.
	ExitData_type_primitive(c *Data_type_primitiveContext)

	// ExitData_type_usual is called when exiting the data_type_usual production.
	ExitData_type_usual(c *Data_type_usualContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitData_type_or_implicit is called when exiting the data_type_or_implicit production.
	ExitData_type_or_implicit(c *Data_type_or_implicitContext)

	// ExitImplicit_data_type is called when exiting the implicit_data_type production.
	ExitImplicit_data_type(c *Implicit_data_typeContext)

	// ExitSequence_list_of_arguments_named_item is called when exiting the sequence_list_of_arguments_named_item production.
	ExitSequence_list_of_arguments_named_item(c *Sequence_list_of_arguments_named_itemContext)

	// ExitSequence_list_of_arguments is called when exiting the sequence_list_of_arguments production.
	ExitSequence_list_of_arguments(c *Sequence_list_of_argumentsContext)

	// ExitList_of_arguments_named_item is called when exiting the list_of_arguments_named_item production.
	ExitList_of_arguments_named_item(c *List_of_arguments_named_itemContext)

	// ExitList_of_arguments is called when exiting the list_of_arguments production.
	ExitList_of_arguments(c *List_of_argumentsContext)

	// ExitPrimary_literal is called when exiting the primary_literal production.
	ExitPrimary_literal(c *Primary_literalContext)

	// ExitType_reference is called when exiting the type_reference production.
	ExitType_reference(c *Type_referenceContext)

	// ExitPackage_scope is called when exiting the package_scope production.
	ExitPackage_scope(c *Package_scopeContext)

	// ExitPs_identifier is called when exiting the ps_identifier production.
	ExitPs_identifier(c *Ps_identifierContext)

	// ExitList_of_parameter_value_assignments is called when exiting the list_of_parameter_value_assignments production.
	ExitList_of_parameter_value_assignments(c *List_of_parameter_value_assignmentsContext)

	// ExitParameter_value_assignment is called when exiting the parameter_value_assignment production.
	ExitParameter_value_assignment(c *Parameter_value_assignmentContext)

	// ExitClass_type is called when exiting the class_type production.
	ExitClass_type(c *Class_typeContext)

	// ExitClass_scope is called when exiting the class_scope production.
	ExitClass_scope(c *Class_scopeContext)

	// ExitRange_expression is called when exiting the range_expression production.
	ExitRange_expression(c *Range_expressionContext)

	// ExitConstant_range_expression is called when exiting the constant_range_expression production.
	ExitConstant_range_expression(c *Constant_range_expressionContext)

	// ExitConstant_mintypmax_expression is called when exiting the constant_mintypmax_expression production.
	ExitConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// ExitMintypmax_expression is called when exiting the mintypmax_expression production.
	ExitMintypmax_expression(c *Mintypmax_expressionContext)

	// ExitNamed_parameter_assignment is called when exiting the named_parameter_assignment production.
	ExitNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// ExitPrimaryLit is called when exiting the PrimaryLit production.
	ExitPrimaryLit(c *PrimaryLitContext)

	// ExitPrimaryRandomize is called when exiting the PrimaryRandomize production.
	ExitPrimaryRandomize(c *PrimaryRandomizeContext)

	// ExitPrimaryAssig is called when exiting the PrimaryAssig production.
	ExitPrimaryAssig(c *PrimaryAssigContext)

	// ExitPrimaryBitSelect is called when exiting the PrimaryBitSelect production.
	ExitPrimaryBitSelect(c *PrimaryBitSelectContext)

	// ExitPrimaryTfCall is called when exiting the PrimaryTfCall production.
	ExitPrimaryTfCall(c *PrimaryTfCallContext)

	// ExitPrimaryTypeRef is called when exiting the PrimaryTypeRef production.
	ExitPrimaryTypeRef(c *PrimaryTypeRefContext)

	// ExitPrimaryCallArrayMethodNoArgs is called when exiting the PrimaryCallArrayMethodNoArgs production.
	ExitPrimaryCallArrayMethodNoArgs(c *PrimaryCallArrayMethodNoArgsContext)

	// ExitPrimaryCast is called when exiting the PrimaryCast production.
	ExitPrimaryCast(c *PrimaryCastContext)

	// ExitPrimaryPar is called when exiting the PrimaryPar production.
	ExitPrimaryPar(c *PrimaryParContext)

	// ExitPrimaryCall is called when exiting the PrimaryCall production.
	ExitPrimaryCall(c *PrimaryCallContext)

	// ExitPrimaryRandomize2 is called when exiting the PrimaryRandomize2 production.
	ExitPrimaryRandomize2(c *PrimaryRandomize2Context)

	// ExitPrimaryDot is called when exiting the PrimaryDot production.
	ExitPrimaryDot(c *PrimaryDotContext)

	// ExitPrimaryStreaming_concatenation is called when exiting the PrimaryStreaming_concatenation production.
	ExitPrimaryStreaming_concatenation(c *PrimaryStreaming_concatenationContext)

	// ExitPrimaryPath is called when exiting the PrimaryPath production.
	ExitPrimaryPath(c *PrimaryPathContext)

	// ExitPrimaryRange is called when exiting the PrimaryRange production.
	ExitPrimaryRange(c *PrimaryRangeContext)

	// ExitPrimaryCallWith is called when exiting the PrimaryCallWith production.
	ExitPrimaryCallWith(c *PrimaryCallWithContext)

	// ExitPrimaryConcat is called when exiting the PrimaryConcat production.
	ExitPrimaryConcat(c *PrimaryConcatContext)

	// ExitPrimaryCast2 is called when exiting the PrimaryCast2 production.
	ExitPrimaryCast2(c *PrimaryCast2Context)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitInc_or_dec_expressionPre is called when exiting the Inc_or_dec_expressionPre production.
	ExitInc_or_dec_expressionPre(c *Inc_or_dec_expressionPreContext)

	// ExitInc_or_dec_expressionPost is called when exiting the Inc_or_dec_expressionPost production.
	ExitInc_or_dec_expressionPost(c *Inc_or_dec_expressionPostContext)

	// ExitExpressionPrimary is called when exiting the ExpressionPrimary production.
	ExitExpressionPrimary(c *ExpressionPrimaryContext)

	// ExitExpressionInside is called when exiting the ExpressionInside production.
	ExitExpressionInside(c *ExpressionInsideContext)

	// ExitExpressionBinOpAnd is called when exiting the ExpressionBinOpAnd production.
	ExitExpressionBinOpAnd(c *ExpressionBinOpAndContext)

	// ExitExpressionBinOpPower is called when exiting the ExpressionBinOpPower production.
	ExitExpressionBinOpPower(c *ExpressionBinOpPowerContext)

	// ExitExpressionBinOpImpl is called when exiting the ExpressionBinOpImpl production.
	ExitExpressionBinOpImpl(c *ExpressionBinOpImplContext)

	// ExitExpressionBinOpEq is called when exiting the ExpressionBinOpEq production.
	ExitExpressionBinOpEq(c *ExpressionBinOpEqContext)

	// ExitExpressionTernary is called when exiting the ExpressionTernary production.
	ExitExpressionTernary(c *ExpressionTernaryContext)

	// ExitExpressionTaggedId is called when exiting the ExpressionTaggedId production.
	ExitExpressionTaggedId(c *ExpressionTaggedIdContext)

	// ExitExpressionUnary is called when exiting the ExpressionUnary production.
	ExitExpressionUnary(c *ExpressionUnaryContext)

	// ExitExpressionAssign is called when exiting the ExpressionAssign production.
	ExitExpressionAssign(c *ExpressionAssignContext)

	// ExitExpressionIncDec is called when exiting the ExpressionIncDec production.
	ExitExpressionIncDec(c *ExpressionIncDecContext)

	// ExitExpressionBinOpMul is called when exiting the ExpressionBinOpMul production.
	ExitExpressionBinOpMul(c *ExpressionBinOpMulContext)

	// ExitExpressionBinOpShift is called when exiting the ExpressionBinOpShift production.
	ExitExpressionBinOpShift(c *ExpressionBinOpShiftContext)

	// ExitExpressionBinOpCmp is called when exiting the ExpressionBinOpCmp production.
	ExitExpressionBinOpCmp(c *ExpressionBinOpCmpContext)

	// ExitExpressionBinOpBitAnd is called when exiting the ExpressionBinOpBitAnd production.
	ExitExpressionBinOpBitAnd(c *ExpressionBinOpBitAndContext)

	// ExitExpressionBinOpAdd is called when exiting the ExpressionBinOpAdd production.
	ExitExpressionBinOpAdd(c *ExpressionBinOpAddContext)

	// ExitExpressionBinOpBitXor is called when exiting the ExpressionBinOpBitXor production.
	ExitExpressionBinOpBitXor(c *ExpressionBinOpBitXorContext)

	// ExitExpressionBinOpBitOr is called when exiting the ExpressionBinOpBitOr production.
	ExitExpressionBinOpBitOr(c *ExpressionBinOpBitOrContext)

	// ExitExpressionBinOpOr is called when exiting the ExpressionBinOpOr production.
	ExitExpressionBinOpOr(c *ExpressionBinOpOrContext)

	// ExitExpressionTripleAnd is called when exiting the ExpressionTripleAnd production.
	ExitExpressionTripleAnd(c *ExpressionTripleAndContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitDynamic_array_new is called when exiting the dynamic_array_new production.
	ExitDynamic_array_new(c *Dynamic_array_newContext)

	// ExitConst_or_range_expression is called when exiting the const_or_range_expression production.
	ExitConst_or_range_expression(c *Const_or_range_expressionContext)

	// ExitVariable_decl_assignment is called when exiting the variable_decl_assignment production.
	ExitVariable_decl_assignment(c *Variable_decl_assignmentContext)

	// ExitAssignment_pattern_variable_lvalue is called when exiting the assignment_pattern_variable_lvalue production.
	ExitAssignment_pattern_variable_lvalue(c *Assignment_pattern_variable_lvalueContext)

	// ExitStream_operator is called when exiting the stream_operator production.
	ExitStream_operator(c *Stream_operatorContext)

	// ExitSlice_size is called when exiting the slice_size production.
	ExitSlice_size(c *Slice_sizeContext)

	// ExitStreaming_concatenation is called when exiting the streaming_concatenation production.
	ExitStreaming_concatenation(c *Streaming_concatenationContext)

	// ExitStream_concatenation is called when exiting the stream_concatenation production.
	ExitStream_concatenation(c *Stream_concatenationContext)

	// ExitStream_expression is called when exiting the stream_expression production.
	ExitStream_expression(c *Stream_expressionContext)

	// ExitArray_range_expression is called when exiting the array_range_expression production.
	ExitArray_range_expression(c *Array_range_expressionContext)

	// ExitOpen_range_list is called when exiting the open_range_list production.
	ExitOpen_range_list(c *Open_range_listContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitAssignment_pattern is called when exiting the assignment_pattern production.
	ExitAssignment_pattern(c *Assignment_patternContext)

	// ExitStructure_pattern_key is called when exiting the structure_pattern_key production.
	ExitStructure_pattern_key(c *Structure_pattern_keyContext)

	// ExitArray_pattern_key is called when exiting the array_pattern_key production.
	ExitArray_pattern_key(c *Array_pattern_keyContext)

	// ExitAssignment_pattern_key is called when exiting the assignment_pattern_key production.
	ExitAssignment_pattern_key(c *Assignment_pattern_keyContext)

	// ExitStruct_union_member is called when exiting the struct_union_member production.
	ExitStruct_union_member(c *Struct_union_memberContext)

	// ExitData_type_or_void is called when exiting the data_type_or_void production.
	ExitData_type_or_void(c *Data_type_or_voidContext)

	// ExitEnum_name_declaration is called when exiting the enum_name_declaration production.
	ExitEnum_name_declaration(c *Enum_name_declarationContext)

	// ExitAssignment_pattern_expression is called when exiting the assignment_pattern_expression production.
	ExitAssignment_pattern_expression(c *Assignment_pattern_expressionContext)

	// ExitAssignment_pattern_expression_type is called when exiting the assignment_pattern_expression_type production.
	ExitAssignment_pattern_expression_type(c *Assignment_pattern_expression_typeContext)

	// ExitNet_lvalue is called when exiting the net_lvalue production.
	ExitNet_lvalue(c *Net_lvalueContext)

	// ExitVarLConcat is called when exiting the VarLConcat production.
	ExitVarLConcat(c *VarLConcatContext)

	// ExitVarLPath is called when exiting the VarLPath production.
	ExitVarLPath(c *VarLPathContext)

	// ExitVarLAssign is called when exiting the VarLAssign production.
	ExitVarLAssign(c *VarLAssignContext)

	// ExitVarLStreamConcat is called when exiting the VarLStreamConcat production.
	ExitVarLStreamConcat(c *VarLStreamConcatContext)

	// ExitSolve_before_list is called when exiting the solve_before_list production.
	ExitSolve_before_list(c *Solve_before_listContext)

	// ExitConstraint_block_item is called when exiting the constraint_block_item production.
	ExitConstraint_block_item(c *Constraint_block_itemContext)

	// ExitConstraint_expression is called when exiting the constraint_expression production.
	ExitConstraint_expression(c *Constraint_expressionContext)

	// ExitUniqueness_constraint is called when exiting the uniqueness_constraint production.
	ExitUniqueness_constraint(c *Uniqueness_constraintContext)

	// ExitConstraint_set is called when exiting the constraint_set production.
	ExitConstraint_set(c *Constraint_setContext)

	// ExitRandomize_call is called when exiting the randomize_call production.
	ExitRandomize_call(c *Randomize_callContext)

	// ExitModule_header_common is called when exiting the module_header_common production.
	ExitModule_header_common(c *Module_header_commonContext)

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitModule_keyword is called when exiting the module_keyword production.
	ExitModule_keyword(c *Module_keywordContext)

	// ExitNet_port_type is called when exiting the net_port_type production.
	ExitNet_port_type(c *Net_port_typeContext)

	// ExitVar_data_type is called when exiting the var_data_type production.
	ExitVar_data_type(c *Var_data_typeContext)

	// ExitNet_or_var_data_type is called when exiting the net_or_var_data_type production.
	ExitNet_or_var_data_type(c *Net_or_var_data_typeContext)

	// ExitList_of_defparam_assignments is called when exiting the list_of_defparam_assignments production.
	ExitList_of_defparam_assignments(c *List_of_defparam_assignmentsContext)

	// ExitList_of_net_decl_assignments is called when exiting the list_of_net_decl_assignments production.
	ExitList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// ExitList_of_specparam_assignments is called when exiting the list_of_specparam_assignments production.
	ExitList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// ExitList_of_variable_decl_assignments is called when exiting the list_of_variable_decl_assignments production.
	ExitList_of_variable_decl_assignments(c *List_of_variable_decl_assignmentsContext)

	// ExitList_of_variable_identifiers_item is called when exiting the list_of_variable_identifiers_item production.
	ExitList_of_variable_identifiers_item(c *List_of_variable_identifiers_itemContext)

	// ExitList_of_variable_identifiers is called when exiting the list_of_variable_identifiers production.
	ExitList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// ExitList_of_variable_port_identifiers is called when exiting the list_of_variable_port_identifiers production.
	ExitList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// ExitDefparam_assignment is called when exiting the defparam_assignment production.
	ExitDefparam_assignment(c *Defparam_assignmentContext)

	// ExitNet_decl_assignment is called when exiting the net_decl_assignment production.
	ExitNet_decl_assignment(c *Net_decl_assignmentContext)

	// ExitSpecparam_assignment is called when exiting the specparam_assignment production.
	ExitSpecparam_assignment(c *Specparam_assignmentContext)

	// ExitError_limit_value is called when exiting the error_limit_value production.
	ExitError_limit_value(c *Error_limit_valueContext)

	// ExitReject_limit_value is called when exiting the reject_limit_value production.
	ExitReject_limit_value(c *Reject_limit_valueContext)

	// ExitPulse_control_specparam is called when exiting the pulse_control_specparam production.
	ExitPulse_control_specparam(c *Pulse_control_specparamContext)

	// ExitIdentifier_doted_index_at_end is called when exiting the identifier_doted_index_at_end production.
	ExitIdentifier_doted_index_at_end(c *Identifier_doted_index_at_endContext)

	// ExitSpecify_terminal_descriptor is called when exiting the specify_terminal_descriptor production.
	ExitSpecify_terminal_descriptor(c *Specify_terminal_descriptorContext)

	// ExitSpecify_input_terminal_descriptor is called when exiting the specify_input_terminal_descriptor production.
	ExitSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// ExitSpecify_output_terminal_descriptor is called when exiting the specify_output_terminal_descriptor production.
	ExitSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// ExitSpecify_item is called when exiting the specify_item production.
	ExitSpecify_item(c *Specify_itemContext)

	// ExitPulsestyle_declaration is called when exiting the pulsestyle_declaration production.
	ExitPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// ExitShowcancelled_declaration is called when exiting the showcancelled_declaration production.
	ExitShowcancelled_declaration(c *Showcancelled_declarationContext)

	// ExitPath_declaration is called when exiting the path_declaration production.
	ExitPath_declaration(c *Path_declarationContext)

	// ExitSimple_path_declaration is called when exiting the simple_path_declaration production.
	ExitSimple_path_declaration(c *Simple_path_declarationContext)

	// ExitPath_delay_value is called when exiting the path_delay_value production.
	ExitPath_delay_value(c *Path_delay_valueContext)

	// ExitList_of_path_outputs is called when exiting the list_of_path_outputs production.
	ExitList_of_path_outputs(c *List_of_path_outputsContext)

	// ExitList_of_path_inputs is called when exiting the list_of_path_inputs production.
	ExitList_of_path_inputs(c *List_of_path_inputsContext)

	// ExitList_of_paths is called when exiting the list_of_paths production.
	ExitList_of_paths(c *List_of_pathsContext)

	// ExitList_of_path_delay_expressions is called when exiting the list_of_path_delay_expressions production.
	ExitList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// ExitT_path_delay_expression is called when exiting the t_path_delay_expression production.
	ExitT_path_delay_expression(c *T_path_delay_expressionContext)

	// ExitTrise_path_delay_expression is called when exiting the trise_path_delay_expression production.
	ExitTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// ExitTfall_path_delay_expression is called when exiting the tfall_path_delay_expression production.
	ExitTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// ExitTz_path_delay_expression is called when exiting the tz_path_delay_expression production.
	ExitTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// ExitT01_path_delay_expression is called when exiting the t01_path_delay_expression production.
	ExitT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// ExitT10_path_delay_expression is called when exiting the t10_path_delay_expression production.
	ExitT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// ExitT0z_path_delay_expression is called when exiting the t0z_path_delay_expression production.
	ExitT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// ExitTz1_path_delay_expression is called when exiting the tz1_path_delay_expression production.
	ExitTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// ExitT1z_path_delay_expression is called when exiting the t1z_path_delay_expression production.
	ExitT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// ExitTz0_path_delay_expression is called when exiting the tz0_path_delay_expression production.
	ExitTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// ExitT0x_path_delay_expression is called when exiting the t0x_path_delay_expression production.
	ExitT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// ExitTx1_path_delay_expression is called when exiting the tx1_path_delay_expression production.
	ExitTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// ExitT1x_path_delay_expression is called when exiting the t1x_path_delay_expression production.
	ExitT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// ExitTx0_path_delay_expression is called when exiting the tx0_path_delay_expression production.
	ExitTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// ExitTxz_path_delay_expression is called when exiting the txz_path_delay_expression production.
	ExitTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// ExitTzx_path_delay_expression is called when exiting the tzx_path_delay_expression production.
	ExitTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// ExitParallel_path_description is called when exiting the parallel_path_description production.
	ExitParallel_path_description(c *Parallel_path_descriptionContext)

	// ExitFull_path_description is called when exiting the full_path_description production.
	ExitFull_path_description(c *Full_path_descriptionContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitSpecparam_declaration is called when exiting the specparam_declaration production.
	ExitSpecparam_declaration(c *Specparam_declarationContext)

	// ExitEdge_sensitive_path_declaration is called when exiting the edge_sensitive_path_declaration production.
	ExitEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// ExitParallel_edge_sensitive_path_description is called when exiting the parallel_edge_sensitive_path_description production.
	ExitParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// ExitFull_edge_sensitive_path_description is called when exiting the full_edge_sensitive_path_description production.
	ExitFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// ExitData_source_expression is called when exiting the data_source_expression production.
	ExitData_source_expression(c *Data_source_expressionContext)

	// ExitData_declaration is called when exiting the data_declaration production.
	ExitData_declaration(c *Data_declarationContext)

	// ExitModule_path_expression is called when exiting the module_path_expression production.
	ExitModule_path_expression(c *Module_path_expressionContext)

	// ExitState_dependent_path_declaration is called when exiting the state_dependent_path_declaration production.
	ExitState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// ExitPackage_export_declaration is called when exiting the package_export_declaration production.
	ExitPackage_export_declaration(c *Package_export_declarationContext)

	// ExitGenvar_declaration is called when exiting the genvar_declaration production.
	ExitGenvar_declaration(c *Genvar_declarationContext)

	// ExitNet_declaration is called when exiting the net_declaration production.
	ExitNet_declaration(c *Net_declarationContext)

	// ExitParameter_port_list is called when exiting the parameter_port_list production.
	ExitParameter_port_list(c *Parameter_port_listContext)

	// ExitParamPortType is called when exiting the ParamPortType production.
	ExitParamPortType(c *ParamPortTypeContext)

	// ExitParamSimple is called when exiting the ParamSimple production.
	ExitParamSimple(c *ParamSimpleContext)

	// ExitParamLocal is called when exiting the ParamLocal production.
	ExitParamLocal(c *ParamLocalContext)

	// ExitParamAssign is called when exiting the ParamAssign production.
	ExitParamAssign(c *ParamAssignContext)

	// ExitList_of_port_declarations_ansi_item is called when exiting the list_of_port_declarations_ansi_item production.
	ExitList_of_port_declarations_ansi_item(c *List_of_port_declarations_ansi_itemContext)

	// ExitList_of_port_declarations is called when exiting the list_of_port_declarations production.
	ExitList_of_port_declarations(c *List_of_port_declarationsContext)

	// ExitNonansi_port_declaration is called when exiting the nonansi_port_declaration production.
	ExitNonansi_port_declaration(c *Nonansi_port_declarationContext)

	// ExitNonansi_port is called when exiting the nonansi_port production.
	ExitNonansi_port(c *Nonansi_portContext)

	// ExitNonansi_port__expr is called when exiting the nonansi_port__expr production.
	ExitNonansi_port__expr(c *Nonansi_port__exprContext)

	// ExitPort_identifier is called when exiting the port_identifier production.
	ExitPort_identifier(c *Port_identifierContext)

	// ExitAnsi_port_declaration is called when exiting the ansi_port_declaration production.
	ExitAnsi_port_declaration(c *Ansi_port_declarationContext)

	// ExitSystem_timing_check is called when exiting the system_timing_check production.
	ExitSystem_timing_check(c *System_timing_checkContext)

	// ExitDolar_setup_timing_check is called when exiting the dolar_setup_timing_check production.
	ExitDolar_setup_timing_check(c *Dolar_setup_timing_checkContext)

	// ExitDolar_hold_timing_check is called when exiting the dolar_hold_timing_check production.
	ExitDolar_hold_timing_check(c *Dolar_hold_timing_checkContext)

	// ExitDolar_setuphold_timing_check is called when exiting the dolar_setuphold_timing_check production.
	ExitDolar_setuphold_timing_check(c *Dolar_setuphold_timing_checkContext)

	// ExitDolar_recovery_timing_check is called when exiting the dolar_recovery_timing_check production.
	ExitDolar_recovery_timing_check(c *Dolar_recovery_timing_checkContext)

	// ExitDolar_removal_timing_check is called when exiting the dolar_removal_timing_check production.
	ExitDolar_removal_timing_check(c *Dolar_removal_timing_checkContext)

	// ExitDolar_recrem_timing_check is called when exiting the dolar_recrem_timing_check production.
	ExitDolar_recrem_timing_check(c *Dolar_recrem_timing_checkContext)

	// ExitDolar_skew_timing_check is called when exiting the dolar_skew_timing_check production.
	ExitDolar_skew_timing_check(c *Dolar_skew_timing_checkContext)

	// ExitDolar_timeskew_timing_check is called when exiting the dolar_timeskew_timing_check production.
	ExitDolar_timeskew_timing_check(c *Dolar_timeskew_timing_checkContext)

	// ExitDolar_fullskew_timing_check is called when exiting the dolar_fullskew_timing_check production.
	ExitDolar_fullskew_timing_check(c *Dolar_fullskew_timing_checkContext)

	// ExitDolar_period_timing_check is called when exiting the dolar_period_timing_check production.
	ExitDolar_period_timing_check(c *Dolar_period_timing_checkContext)

	// ExitDolar_width_timing_check is called when exiting the dolar_width_timing_check production.
	ExitDolar_width_timing_check(c *Dolar_width_timing_checkContext)

	// ExitDolar_nochange_timing_check is called when exiting the dolar_nochange_timing_check production.
	ExitDolar_nochange_timing_check(c *Dolar_nochange_timing_checkContext)

	// ExitTimecheck_condition is called when exiting the timecheck_condition production.
	ExitTimecheck_condition(c *Timecheck_conditionContext)

	// ExitControlled_reference_event is called when exiting the controlled_reference_event production.
	ExitControlled_reference_event(c *Controlled_reference_eventContext)

	// ExitDelayed_reference is called when exiting the delayed_reference production.
	ExitDelayed_reference(c *Delayed_referenceContext)

	// ExitEnd_edge_offset is called when exiting the end_edge_offset production.
	ExitEnd_edge_offset(c *End_edge_offsetContext)

	// ExitEvent_based_flag is called when exiting the event_based_flag production.
	ExitEvent_based_flag(c *Event_based_flagContext)

	// ExitNotifier is called when exiting the notifier production.
	ExitNotifier(c *NotifierContext)

	// ExitRemain_active_flag is called when exiting the remain_active_flag production.
	ExitRemain_active_flag(c *Remain_active_flagContext)

	// ExitTimestamp_condition is called when exiting the timestamp_condition production.
	ExitTimestamp_condition(c *Timestamp_conditionContext)

	// ExitStart_edge_offset is called when exiting the start_edge_offset production.
	ExitStart_edge_offset(c *Start_edge_offsetContext)

	// ExitThreshold is called when exiting the threshold production.
	ExitThreshold(c *ThresholdContext)

	// ExitTiming_check_limit is called when exiting the timing_check_limit production.
	ExitTiming_check_limit(c *Timing_check_limitContext)

	// ExitTiming_check_event is called when exiting the timing_check_event production.
	ExitTiming_check_event(c *Timing_check_eventContext)

	// ExitTiming_check_condition is called when exiting the timing_check_condition production.
	ExitTiming_check_condition(c *Timing_check_conditionContext)

	// ExitScalar_timing_check_condition is called when exiting the scalar_timing_check_condition production.
	ExitScalar_timing_check_condition(c *Scalar_timing_check_conditionContext)

	// ExitControlled_timing_check_event is called when exiting the controlled_timing_check_event production.
	ExitControlled_timing_check_event(c *Controlled_timing_check_eventContext)

	// ExitFunction_data_type_or_implicit is called when exiting the function_data_type_or_implicit production.
	ExitFunction_data_type_or_implicit(c *Function_data_type_or_implicitContext)

	// ExitExtern_tf_declaration is called when exiting the extern_tf_declaration production.
	ExitExtern_tf_declaration(c *Extern_tf_declarationContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitTask_prototype is called when exiting the task_prototype production.
	ExitTask_prototype(c *Task_prototypeContext)

	// ExitFunction_prototype is called when exiting the function_prototype production.
	ExitFunction_prototype(c *Function_prototypeContext)

	// ExitDpi_import_export is called when exiting the dpi_import_export production.
	ExitDpi_import_export(c *Dpi_import_exportContext)

	// ExitDpi_function_import_property is called when exiting the dpi_function_import_property production.
	ExitDpi_function_import_property(c *Dpi_function_import_propertyContext)

	// ExitDpi_task_import_property is called when exiting the dpi_task_import_property production.
	ExitDpi_task_import_property(c *Dpi_task_import_propertyContext)

	// ExitTask_and_function_declaration_common is called when exiting the task_and_function_declaration_common production.
	ExitTask_and_function_declaration_common(c *Task_and_function_declaration_commonContext)

	// ExitTask_declaration is called when exiting the task_declaration production.
	ExitTask_declaration(c *Task_declarationContext)

	// ExitMethod_prototype is called when exiting the method_prototype production.
	ExitMethod_prototype(c *Method_prototypeContext)

	// ExitExtern_constraint_declaration is called when exiting the extern_constraint_declaration production.
	ExitExtern_constraint_declaration(c *Extern_constraint_declarationContext)

	// ExitConstraint_block is called when exiting the constraint_block production.
	ExitConstraint_block(c *Constraint_blockContext)

	// ExitChecker_port_list is called when exiting the checker_port_list production.
	ExitChecker_port_list(c *Checker_port_listContext)

	// ExitChecker_port_item is called when exiting the checker_port_item production.
	ExitChecker_port_item(c *Checker_port_itemContext)

	// ExitChecker_port_direction is called when exiting the checker_port_direction production.
	ExitChecker_port_direction(c *Checker_port_directionContext)

	// ExitChecker_declaration is called when exiting the checker_declaration production.
	ExitChecker_declaration(c *Checker_declarationContext)

	// ExitClass_declaration is called when exiting the class_declaration production.
	ExitClass_declaration(c *Class_declarationContext)

	// ExitAlways_construct is called when exiting the always_construct production.
	ExitAlways_construct(c *Always_constructContext)

	// ExitInterface_class_type is called when exiting the interface_class_type production.
	ExitInterface_class_type(c *Interface_class_typeContext)

	// ExitInterface_class_declaration is called when exiting the interface_class_declaration production.
	ExitInterface_class_declaration(c *Interface_class_declarationContext)

	// ExitInterface_class_item is called when exiting the interface_class_item production.
	ExitInterface_class_item(c *Interface_class_itemContext)

	// ExitInterface_class_method is called when exiting the interface_class_method production.
	ExitInterface_class_method(c *Interface_class_methodContext)

	// ExitPackage_declaration is called when exiting the package_declaration production.
	ExitPackage_declaration(c *Package_declarationContext)

	// ExitPackage_item is called when exiting the package_item production.
	ExitPackage_item(c *Package_itemContext)

	// ExitPackage_item_item is called when exiting the package_item_item production.
	ExitPackage_item_item(c *Package_item_itemContext)

	// ExitProgram_declaration is called when exiting the program_declaration production.
	ExitProgram_declaration(c *Program_declarationContext)

	// ExitProgram_header is called when exiting the program_header production.
	ExitProgram_header(c *Program_headerContext)

	// ExitProgram_item is called when exiting the program_item production.
	ExitProgram_item(c *Program_itemContext)

	// ExitNon_port_program_item is called when exiting the non_port_program_item production.
	ExitNon_port_program_item(c *Non_port_program_itemContext)

	// ExitAnonymous_program is called when exiting the anonymous_program production.
	ExitAnonymous_program(c *Anonymous_programContext)

	// ExitAnonymous_program_item is called when exiting the anonymous_program_item production.
	ExitAnonymous_program_item(c *Anonymous_program_itemContext)

	// ExitSequence_declaration is called when exiting the sequence_declaration production.
	ExitSequence_declaration(c *Sequence_declarationContext)

	// ExitSequence_port_list is called when exiting the sequence_port_list production.
	ExitSequence_port_list(c *Sequence_port_listContext)

	// ExitSequence_port_item is called when exiting the sequence_port_item production.
	ExitSequence_port_item(c *Sequence_port_itemContext)

	// ExitProperty_declaration is called when exiting the property_declaration production.
	ExitProperty_declaration(c *Property_declarationContext)

	// ExitProperty_port_list is called when exiting the property_port_list production.
	ExitProperty_port_list(c *Property_port_listContext)

	// ExitProperty_port_item is called when exiting the property_port_item production.
	ExitProperty_port_item(c *Property_port_itemContext)

	// ExitContinuous_assign is called when exiting the continuous_assign production.
	ExitContinuous_assign(c *Continuous_assignContext)

	// ExitChecker_or_generate_item is called when exiting the checker_or_generate_item production.
	ExitChecker_or_generate_item(c *Checker_or_generate_itemContext)

	// ExitConstraint_prototype is called when exiting the constraint_prototype production.
	ExitConstraint_prototype(c *Constraint_prototypeContext)

	// ExitClass_constraint is called when exiting the class_constraint production.
	ExitClass_constraint(c *Class_constraintContext)

	// ExitConstraint_declaration is called when exiting the constraint_declaration production.
	ExitConstraint_declaration(c *Constraint_declarationContext)

	// ExitClass_constructor_declaration is called when exiting the class_constructor_declaration production.
	ExitClass_constructor_declaration(c *Class_constructor_declarationContext)

	// ExitClass_property is called when exiting the class_property production.
	ExitClass_property(c *Class_propertyContext)

	// ExitClass_method is called when exiting the class_method production.
	ExitClass_method(c *Class_methodContext)

	// ExitClass_constructor_prototype is called when exiting the class_constructor_prototype production.
	ExitClass_constructor_prototype(c *Class_constructor_prototypeContext)

	// ExitClass_item is called when exiting the class_item production.
	ExitClass_item(c *Class_itemContext)

	// ExitParameter_override is called when exiting the parameter_override production.
	ExitParameter_override(c *Parameter_overrideContext)

	// ExitGate_instantiation is called when exiting the gate_instantiation production.
	ExitGate_instantiation(c *Gate_instantiationContext)

	// ExitEnable_gate_or_mos_switch_or_cmos_switch_instance is called when exiting the enable_gate_or_mos_switch_or_cmos_switch_instance production.
	ExitEnable_gate_or_mos_switch_or_cmos_switch_instance(c *Enable_gate_or_mos_switch_or_cmos_switch_instanceContext)

	// ExitN_input_gate_instance is called when exiting the n_input_gate_instance production.
	ExitN_input_gate_instance(c *N_input_gate_instanceContext)

	// ExitN_output_gate_instance is called when exiting the n_output_gate_instance production.
	ExitN_output_gate_instance(c *N_output_gate_instanceContext)

	// ExitPass_switch_instance is called when exiting the pass_switch_instance production.
	ExitPass_switch_instance(c *Pass_switch_instanceContext)

	// ExitPass_enable_switch_instance is called when exiting the pass_enable_switch_instance production.
	ExitPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// ExitPull_gate_instance is called when exiting the pull_gate_instance production.
	ExitPull_gate_instance(c *Pull_gate_instanceContext)

	// ExitPulldown_strength is called when exiting the pulldown_strength production.
	ExitPulldown_strength(c *Pulldown_strengthContext)

	// ExitPullup_strength is called when exiting the pullup_strength production.
	ExitPullup_strength(c *Pullup_strengthContext)

	// ExitEnable_terminal is called when exiting the enable_terminal production.
	ExitEnable_terminal(c *Enable_terminalContext)

	// ExitInout_terminal is called when exiting the inout_terminal production.
	ExitInout_terminal(c *Inout_terminalContext)

	// ExitInput_terminal is called when exiting the input_terminal production.
	ExitInput_terminal(c *Input_terminalContext)

	// ExitOutput_terminal is called when exiting the output_terminal production.
	ExitOutput_terminal(c *Output_terminalContext)

	// ExitUdp_instantiation is called when exiting the udp_instantiation production.
	ExitUdp_instantiation(c *Udp_instantiationContext)

	// ExitUdp_instance is called when exiting the udp_instance production.
	ExitUdp_instance(c *Udp_instanceContext)

	// ExitUdp_instance_body is called when exiting the udp_instance_body production.
	ExitUdp_instance_body(c *Udp_instance_bodyContext)

	// ExitModule_or_interface_or_program_or_udp_instantiation is called when exiting the module_or_interface_or_program_or_udp_instantiation production.
	ExitModule_or_interface_or_program_or_udp_instantiation(c *Module_or_interface_or_program_or_udp_instantiationContext)

	// ExitHierarchical_instance is called when exiting the hierarchical_instance production.
	ExitHierarchical_instance(c *Hierarchical_instanceContext)

	// ExitList_of_port_connections is called when exiting the list_of_port_connections production.
	ExitList_of_port_connections(c *List_of_port_connectionsContext)

	// ExitOrdered_port_connection is called when exiting the ordered_port_connection production.
	ExitOrdered_port_connection(c *Ordered_port_connectionContext)

	// ExitNamed_port_connection is called when exiting the named_port_connection production.
	ExitNamed_port_connection(c *Named_port_connectionContext)

	// ExitBind_directive is called when exiting the bind_directive production.
	ExitBind_directive(c *Bind_directiveContext)

	// ExitBind_target_instance is called when exiting the bind_target_instance production.
	ExitBind_target_instance(c *Bind_target_instanceContext)

	// ExitBind_target_instance_list is called when exiting the bind_target_instance_list production.
	ExitBind_target_instance_list(c *Bind_target_instance_listContext)

	// ExitBind_instantiation is called when exiting the bind_instantiation production.
	ExitBind_instantiation(c *Bind_instantiationContext)

	// ExitConfig_declaration is called when exiting the config_declaration production.
	ExitConfig_declaration(c *Config_declarationContext)

	// ExitDesign_statement is called when exiting the design_statement production.
	ExitDesign_statement(c *Design_statementContext)

	// ExitConfig_rule_statement is called when exiting the config_rule_statement production.
	ExitConfig_rule_statement(c *Config_rule_statementContext)

	// ExitInst_clause is called when exiting the inst_clause production.
	ExitInst_clause(c *Inst_clauseContext)

	// ExitInst_name is called when exiting the inst_name production.
	ExitInst_name(c *Inst_nameContext)

	// ExitCell_clause is called when exiting the cell_clause production.
	ExitCell_clause(c *Cell_clauseContext)

	// ExitLiblist_clause is called when exiting the liblist_clause production.
	ExitLiblist_clause(c *Liblist_clauseContext)

	// ExitUse_clause is called when exiting the use_clause production.
	ExitUse_clause(c *Use_clauseContext)

	// ExitNet_alias is called when exiting the net_alias production.
	ExitNet_alias(c *Net_aliasContext)

	// ExitSpecify_block is called when exiting the specify_block production.
	ExitSpecify_block(c *Specify_blockContext)

	// ExitGenerate_region is called when exiting the generate_region production.
	ExitGenerate_region(c *Generate_regionContext)

	// ExitGenvar_expression is called when exiting the genvar_expression production.
	ExitGenvar_expression(c *Genvar_expressionContext)

	// ExitLoop_generate_construct is called when exiting the loop_generate_construct production.
	ExitLoop_generate_construct(c *Loop_generate_constructContext)

	// ExitGenvar_initialization is called when exiting the genvar_initialization production.
	ExitGenvar_initialization(c *Genvar_initializationContext)

	// ExitGenIterPostfix is called when exiting the GenIterPostfix production.
	ExitGenIterPostfix(c *GenIterPostfixContext)

	// ExitGenIterPrefix is called when exiting the GenIterPrefix production.
	ExitGenIterPrefix(c *GenIterPrefixContext)

	// ExitConditional_generate_construct is called when exiting the conditional_generate_construct production.
	ExitConditional_generate_construct(c *Conditional_generate_constructContext)

	// ExitIf_generate_construct is called when exiting the if_generate_construct production.
	ExitIf_generate_construct(c *If_generate_constructContext)

	// ExitCase_generate_construct is called when exiting the case_generate_construct production.
	ExitCase_generate_construct(c *Case_generate_constructContext)

	// ExitCase_generate_item is called when exiting the case_generate_item production.
	ExitCase_generate_item(c *Case_generate_itemContext)

	// ExitGenerate_begin_end_block is called when exiting the generate_begin_end_block production.
	ExitGenerate_begin_end_block(c *Generate_begin_end_blockContext)

	// ExitGenerate_item_item is called when exiting the generate_item_item production.
	ExitGenerate_item_item(c *Generate_item_itemContext)

	// ExitGenItemItem is called when exiting the GenItemItem production.
	ExitGenItemItem(c *GenItemItemContext)

	// ExitGenItemDataDecl is called when exiting the GenItemDataDecl production.
	ExitGenItemDataDecl(c *GenItemDataDeclContext)

	// ExitGenItemGenReg is called when exiting the GenItemGenReg production.
	ExitGenItemGenReg(c *GenItemGenRegContext)

	// ExitGenItemGenBlock is called when exiting the GenItemGenBlock production.
	ExitGenItemGenBlock(c *GenItemGenBlockContext)

	// ExitProgram_generate_item is called when exiting the program_generate_item production.
	ExitProgram_generate_item(c *Program_generate_itemContext)

	// ExitElaboration_system_task is called when exiting the elaboration_system_task production.
	ExitElaboration_system_task(c *Elaboration_system_taskContext)

	// ExitCoreItemParamOver is called when exiting the CoreItemParamOver production.
	ExitCoreItemParamOver(c *CoreItemParamOverContext)

	// ExitCoreItemGate is called when exiting the CoreItemGate production.
	ExitCoreItemGate(c *CoreItemGateContext)

	// ExitCoreItemUdp is called when exiting the CoreItemUdp production.
	ExitCoreItemUdp(c *CoreItemUdpContext)

	// ExitCoreItemInstance is called when exiting the CoreItemInstance production.
	ExitCoreItemInstance(c *CoreItemInstanceContext)

	// ExitCoreItemParam is called when exiting the CoreItemParam production.
	ExitCoreItemParam(c *CoreItemParamContext)

	// ExitCoreItemNet is called when exiting the CoreItemNet production.
	ExitCoreItemNet(c *CoreItemNetContext)

	// ExitCoreItemData is called when exiting the CoreItemData production.
	ExitCoreItemData(c *CoreItemDataContext)

	// ExitCoreItemTask is called when exiting the CoreItemTask production.
	ExitCoreItemTask(c *CoreItemTaskContext)

	// ExitCoreItemFunction is called when exiting the CoreItemFunction production.
	ExitCoreItemFunction(c *CoreItemFunctionContext)

	// ExitCoreItemChecker is called when exiting the CoreItemChecker production.
	ExitCoreItemChecker(c *CoreItemCheckerContext)

	// ExitCoreItemDPI is called when exiting the CoreItemDPI production.
	ExitCoreItemDPI(c *CoreItemDPIContext)

	// ExitCoreItemExtern is called when exiting the CoreItemExtern production.
	ExitCoreItemExtern(c *CoreItemExternContext)

	// ExitCoreItemClass is called when exiting the CoreItemClass production.
	ExitCoreItemClass(c *CoreItemClassContext)

	// ExitCoreItemIntf is called when exiting the CoreItemIntf production.
	ExitCoreItemIntf(c *CoreItemIntfContext)

	// ExitCoreItemClassCons is called when exiting the CoreItemClassCons production.
	ExitCoreItemClassCons(c *CoreItemClassConsContext)

	// ExitCoreItemCover is called when exiting the CoreItemCover production.
	ExitCoreItemCover(c *CoreItemCoverContext)

	// ExitCoreItemProperty is called when exiting the CoreItemProperty production.
	ExitCoreItemProperty(c *CoreItemPropertyContext)

	// ExitCoreItemSeq is called when exiting the CoreItemSeq production.
	ExitCoreItemSeq(c *CoreItemSeqContext)

	// ExitCoreItemLet is called when exiting the CoreItemLet production.
	ExitCoreItemLet(c *CoreItemLetContext)

	// ExitCoreItemGenVar is called when exiting the CoreItemGenVar production.
	ExitCoreItemGenVar(c *CoreItemGenVarContext)

	// ExitCoreItemClock is called when exiting the CoreItemClock production.
	ExitCoreItemClock(c *CoreItemClockContext)

	// ExitCoreItemAssert is called when exiting the CoreItemAssert production.
	ExitCoreItemAssert(c *CoreItemAssertContext)

	// ExitCoreItemBind is called when exiting the CoreItemBind production.
	ExitCoreItemBind(c *CoreItemBindContext)

	// ExitCoreItemContinuous is called when exiting the CoreItemContinuous production.
	ExitCoreItemContinuous(c *CoreItemContinuousContext)

	// ExitCoreItemNetAlias is called when exiting the CoreItemNetAlias production.
	ExitCoreItemNetAlias(c *CoreItemNetAliasContext)

	// ExitCoreItemInitial is called when exiting the CoreItemInitial production.
	ExitCoreItemInitial(c *CoreItemInitialContext)

	// ExitCoreItemFinal is called when exiting the CoreItemFinal production.
	ExitCoreItemFinal(c *CoreItemFinalContext)

	// ExitCoreItemAlways is called when exiting the CoreItemAlways production.
	ExitCoreItemAlways(c *CoreItemAlwaysContext)

	// ExitCoreItemLoop is called when exiting the CoreItemLoop production.
	ExitCoreItemLoop(c *CoreItemLoopContext)

	// ExitCoreItemCondGen is called when exiting the CoreItemCondGen production.
	ExitCoreItemCondGen(c *CoreItemCondGenContext)

	// ExitCoreItemElab is called when exiting the CoreItemElab production.
	ExitCoreItemElab(c *CoreItemElabContext)

	// ExitModule_item_item is called when exiting the module_item_item production.
	ExitModule_item_item(c *Module_item_itemContext)

	// ExitModuleGenReg is called when exiting the ModuleGenReg production.
	ExitModuleGenReg(c *ModuleGenRegContext)

	// ExitModuleItemItem is called when exiting the ModuleItemItem production.
	ExitModuleItemItem(c *ModuleItemItemContext)

	// ExitModuleSpecBlock is called when exiting the ModuleSpecBlock production.
	ExitModuleSpecBlock(c *ModuleSpecBlockContext)

	// ExitModuleProgDecl is called when exiting the ModuleProgDecl production.
	ExitModuleProgDecl(c *ModuleProgDeclContext)

	// ExitModuleDecl is called when exiting the ModuleDecl production.
	ExitModuleDecl(c *ModuleDeclContext)

	// ExitModuleIntfDecl is called when exiting the ModuleIntfDecl production.
	ExitModuleIntfDecl(c *ModuleIntfDeclContext)

	// ExitModuleTimeDecl is called when exiting the ModuleTimeDecl production.
	ExitModuleTimeDecl(c *ModuleTimeDeclContext)

	// ExitModulePortDecl is called when exiting the ModulePortDecl production.
	ExitModulePortDecl(c *ModulePortDeclContext)
}

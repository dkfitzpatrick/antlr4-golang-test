// Code generated from CPP14Parser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package cpp14 // CPP14Parser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCPP14ParserListener is a complete listener for a parse tree produced by CPP14Parser.
type BaseCPP14ParserListener struct{}

var _ CPP14ParserListener = &BaseCPP14ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCPP14ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCPP14ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCPP14ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCPP14ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *BaseCPP14ParserListener) EnterTranslationUnit(ctx *TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *BaseCPP14ParserListener) ExitTranslationUnit(ctx *TranslationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseCPP14ParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseCPP14ParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterIdExpression is called when production idExpression is entered.
func (s *BaseCPP14ParserListener) EnterIdExpression(ctx *IdExpressionContext) {}

// ExitIdExpression is called when production idExpression is exited.
func (s *BaseCPP14ParserListener) ExitIdExpression(ctx *IdExpressionContext) {}

// EnterUnqualifiedId is called when production unqualifiedId is entered.
func (s *BaseCPP14ParserListener) EnterUnqualifiedId(ctx *UnqualifiedIdContext) {}

// ExitUnqualifiedId is called when production unqualifiedId is exited.
func (s *BaseCPP14ParserListener) ExitUnqualifiedId(ctx *UnqualifiedIdContext) {}

// EnterQualifiedId is called when production qualifiedId is entered.
func (s *BaseCPP14ParserListener) EnterQualifiedId(ctx *QualifiedIdContext) {}

// ExitQualifiedId is called when production qualifiedId is exited.
func (s *BaseCPP14ParserListener) ExitQualifiedId(ctx *QualifiedIdContext) {}

// EnterNestedNameSpecifier is called when production nestedNameSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterNestedNameSpecifier(ctx *NestedNameSpecifierContext) {}

// ExitNestedNameSpecifier is called when production nestedNameSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitNestedNameSpecifier(ctx *NestedNameSpecifierContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseCPP14ParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseCPP14ParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaIntroducer is called when production lambdaIntroducer is entered.
func (s *BaseCPP14ParserListener) EnterLambdaIntroducer(ctx *LambdaIntroducerContext) {}

// ExitLambdaIntroducer is called when production lambdaIntroducer is exited.
func (s *BaseCPP14ParserListener) ExitLambdaIntroducer(ctx *LambdaIntroducerContext) {}

// EnterLambdaCapture is called when production lambdaCapture is entered.
func (s *BaseCPP14ParserListener) EnterLambdaCapture(ctx *LambdaCaptureContext) {}

// ExitLambdaCapture is called when production lambdaCapture is exited.
func (s *BaseCPP14ParserListener) ExitLambdaCapture(ctx *LambdaCaptureContext) {}

// EnterCaptureDefault is called when production captureDefault is entered.
func (s *BaseCPP14ParserListener) EnterCaptureDefault(ctx *CaptureDefaultContext) {}

// ExitCaptureDefault is called when production captureDefault is exited.
func (s *BaseCPP14ParserListener) ExitCaptureDefault(ctx *CaptureDefaultContext) {}

// EnterCaptureList is called when production captureList is entered.
func (s *BaseCPP14ParserListener) EnterCaptureList(ctx *CaptureListContext) {}

// ExitCaptureList is called when production captureList is exited.
func (s *BaseCPP14ParserListener) ExitCaptureList(ctx *CaptureListContext) {}

// EnterCapture is called when production capture is entered.
func (s *BaseCPP14ParserListener) EnterCapture(ctx *CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *BaseCPP14ParserListener) ExitCapture(ctx *CaptureContext) {}

// EnterSimpleCapture is called when production simpleCapture is entered.
func (s *BaseCPP14ParserListener) EnterSimpleCapture(ctx *SimpleCaptureContext) {}

// ExitSimpleCapture is called when production simpleCapture is exited.
func (s *BaseCPP14ParserListener) ExitSimpleCapture(ctx *SimpleCaptureContext) {}

// EnterInitcapture is called when production initcapture is entered.
func (s *BaseCPP14ParserListener) EnterInitcapture(ctx *InitcaptureContext) {}

// ExitInitcapture is called when production initcapture is exited.
func (s *BaseCPP14ParserListener) ExitInitcapture(ctx *InitcaptureContext) {}

// EnterLambdaDeclarator is called when production lambdaDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterLambdaDeclarator(ctx *LambdaDeclaratorContext) {}

// ExitLambdaDeclarator is called when production lambdaDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitLambdaDeclarator(ctx *LambdaDeclaratorContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseCPP14ParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseCPP14ParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is entered.
func (s *BaseCPP14ParserListener) EnterTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) {}

// ExitTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is exited.
func (s *BaseCPP14ParserListener) ExitTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseCPP14ParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseCPP14ParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPseudoDestructorName is called when production pseudoDestructorName is entered.
func (s *BaseCPP14ParserListener) EnterPseudoDestructorName(ctx *PseudoDestructorNameContext) {}

// ExitPseudoDestructorName is called when production pseudoDestructorName is exited.
func (s *BaseCPP14ParserListener) ExitPseudoDestructorName(ctx *PseudoDestructorNameContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseCPP14ParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseCPP14ParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseCPP14ParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseCPP14ParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterNewExpression is called when production newExpression is entered.
func (s *BaseCPP14ParserListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production newExpression is exited.
func (s *BaseCPP14ParserListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterNewPlacement is called when production newPlacement is entered.
func (s *BaseCPP14ParserListener) EnterNewPlacement(ctx *NewPlacementContext) {}

// ExitNewPlacement is called when production newPlacement is exited.
func (s *BaseCPP14ParserListener) ExitNewPlacement(ctx *NewPlacementContext) {}

// EnterNewTypeId is called when production newTypeId is entered.
func (s *BaseCPP14ParserListener) EnterNewTypeId(ctx *NewTypeIdContext) {}

// ExitNewTypeId is called when production newTypeId is exited.
func (s *BaseCPP14ParserListener) ExitNewTypeId(ctx *NewTypeIdContext) {}

// EnterNewDeclarator is called when production newDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterNewDeclarator(ctx *NewDeclaratorContext) {}

// ExitNewDeclarator is called when production newDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitNewDeclarator(ctx *NewDeclaratorContext) {}

// EnterNoPointerNewDeclarator is called when production noPointerNewDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) {}

// ExitNoPointerNewDeclarator is called when production noPointerNewDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) {}

// EnterNewInitializer is called when production newInitializer is entered.
func (s *BaseCPP14ParserListener) EnterNewInitializer(ctx *NewInitializerContext) {}

// ExitNewInitializer is called when production newInitializer is exited.
func (s *BaseCPP14ParserListener) ExitNewInitializer(ctx *NewInitializerContext) {}

// EnterDeleteExpression is called when production deleteExpression is entered.
func (s *BaseCPP14ParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production deleteExpression is exited.
func (s *BaseCPP14ParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterNoExceptExpression is called when production noExceptExpression is entered.
func (s *BaseCPP14ParserListener) EnterNoExceptExpression(ctx *NoExceptExpressionContext) {}

// ExitNoExceptExpression is called when production noExceptExpression is exited.
func (s *BaseCPP14ParserListener) ExitNoExceptExpression(ctx *NoExceptExpressionContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseCPP14ParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseCPP14ParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterPointerMemberExpression is called when production pointerMemberExpression is entered.
func (s *BaseCPP14ParserListener) EnterPointerMemberExpression(ctx *PointerMemberExpressionContext) {}

// ExitPointerMemberExpression is called when production pointerMemberExpression is exited.
func (s *BaseCPP14ParserListener) ExitPointerMemberExpression(ctx *PointerMemberExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseCPP14ParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseCPP14ParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseCPP14ParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseCPP14ParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseCPP14ParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseCPP14ParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *BaseCPP14ParserListener) EnterShiftOperator(ctx *ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *BaseCPP14ParserListener) ExitShiftOperator(ctx *ShiftOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseCPP14ParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseCPP14ParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseCPP14ParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseCPP14ParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseCPP14ParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseCPP14ParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseCPP14ParserListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseCPP14ParserListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseCPP14ParserListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseCPP14ParserListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseCPP14ParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseCPP14ParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseCPP14ParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseCPP14ParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseCPP14ParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseCPP14ParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseCPP14ParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseCPP14ParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseCPP14ParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseCPP14ParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCPP14ParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCPP14ParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseCPP14ParserListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseCPP14ParserListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCPP14ParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCPP14ParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseCPP14ParserListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseCPP14ParserListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseCPP14ParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseCPP14ParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseCPP14ParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseCPP14ParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterStatementSeq is called when production statementSeq is entered.
func (s *BaseCPP14ParserListener) EnterStatementSeq(ctx *StatementSeqContext) {}

// ExitStatementSeq is called when production statementSeq is exited.
func (s *BaseCPP14ParserListener) ExitStatementSeq(ctx *StatementSeqContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseCPP14ParserListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseCPP14ParserListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseCPP14ParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseCPP14ParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseCPP14ParserListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseCPP14ParserListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterForInitStatement is called when production forInitStatement is entered.
func (s *BaseCPP14ParserListener) EnterForInitStatement(ctx *ForInitStatementContext) {}

// ExitForInitStatement is called when production forInitStatement is exited.
func (s *BaseCPP14ParserListener) ExitForInitStatement(ctx *ForInitStatementContext) {}

// EnterForRangeDeclaration is called when production forRangeDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterForRangeDeclaration(ctx *ForRangeDeclarationContext) {}

// ExitForRangeDeclaration is called when production forRangeDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitForRangeDeclaration(ctx *ForRangeDeclarationContext) {}

// EnterForRangeInitializer is called when production forRangeInitializer is entered.
func (s *BaseCPP14ParserListener) EnterForRangeInitializer(ctx *ForRangeInitializerContext) {}

// ExitForRangeInitializer is called when production forRangeInitializer is exited.
func (s *BaseCPP14ParserListener) ExitForRangeInitializer(ctx *ForRangeInitializerContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BaseCPP14ParserListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BaseCPP14ParserListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *BaseCPP14ParserListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *BaseCPP14ParserListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterDeclarationseq is called when production declarationseq is entered.
func (s *BaseCPP14ParserListener) EnterDeclarationseq(ctx *DeclarationseqContext) {}

// ExitDeclarationseq is called when production declarationseq is exited.
func (s *BaseCPP14ParserListener) ExitDeclarationseq(ctx *DeclarationseqContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCPP14ParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCPP14ParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterBlockDeclaration is called when production blockDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterBlockDeclaration(ctx *BlockDeclarationContext) {}

// ExitBlockDeclaration is called when production blockDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitBlockDeclaration(ctx *BlockDeclarationContext) {}

// EnterAliasDeclaration is called when production aliasDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterAliasDeclaration(ctx *AliasDeclarationContext) {}

// ExitAliasDeclaration is called when production aliasDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitAliasDeclaration(ctx *AliasDeclarationContext) {}

// EnterSimpleDeclaration is called when production simpleDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// ExitSimpleDeclaration is called when production simpleDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// EnterEmptyDeclaration is called when production emptyDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterEmptyDeclaration(ctx *EmptyDeclarationContext) {}

// ExitEmptyDeclaration is called when production emptyDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitEmptyDeclaration(ctx *EmptyDeclarationContext) {}

// EnterAttributeDeclaration is called when production attributeDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterAttributeDeclaration(ctx *AttributeDeclarationContext) {}

// ExitAttributeDeclaration is called when production attributeDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitAttributeDeclaration(ctx *AttributeDeclarationContext) {}

// EnterDeclSpecifier is called when production declSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterDeclSpecifier(ctx *DeclSpecifierContext) {}

// ExitDeclSpecifier is called when production declSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitDeclSpecifier(ctx *DeclSpecifierContext) {}

// EnterDeclSpecifierSeq is called when production declSpecifierSeq is entered.
func (s *BaseCPP14ParserListener) EnterDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) {}

// ExitDeclSpecifierSeq is called when production declSpecifierSeq is exited.
func (s *BaseCPP14ParserListener) ExitDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *BaseCPP14ParserListener) EnterTypedefName(ctx *TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *BaseCPP14ParserListener) ExitTypedefName(ctx *TypedefNameContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterTrailingTypeSpecifier is called when production trailingTypeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) {}

// ExitTrailingTypeSpecifier is called when production trailingTypeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) {}

// EnterTypeSpecifierSeq is called when production typeSpecifierSeq is entered.
func (s *BaseCPP14ParserListener) EnterTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) {}

// ExitTypeSpecifierSeq is called when production typeSpecifierSeq is exited.
func (s *BaseCPP14ParserListener) ExitTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) {}

// EnterTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is entered.
func (s *BaseCPP14ParserListener) EnterTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) {
}

// ExitTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is exited.
func (s *BaseCPP14ParserListener) ExitTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) {
}

// EnterSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is entered.
func (s *BaseCPP14ParserListener) EnterSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) {
}

// ExitSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is exited.
func (s *BaseCPP14ParserListener) ExitSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) {
}

// EnterSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is entered.
func (s *BaseCPP14ParserListener) EnterSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) {
}

// ExitSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is exited.
func (s *BaseCPP14ParserListener) ExitSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) {
}

// EnterSimpleTypeSpecifier is called when production simpleTypeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// ExitSimpleTypeSpecifier is called when production simpleTypeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// EnterTheTypeName is called when production theTypeName is entered.
func (s *BaseCPP14ParserListener) EnterTheTypeName(ctx *TheTypeNameContext) {}

// ExitTheTypeName is called when production theTypeName is exited.
func (s *BaseCPP14ParserListener) ExitTheTypeName(ctx *TheTypeNameContext) {}

// EnterDecltypeSpecifier is called when production decltypeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterDecltypeSpecifier(ctx *DecltypeSpecifierContext) {}

// ExitDecltypeSpecifier is called when production decltypeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitDecltypeSpecifier(ctx *DecltypeSpecifierContext) {}

// EnterElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) {}

// ExitElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseCPP14ParserListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseCPP14ParserListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {}

// EnterEnumHead is called when production enumHead is entered.
func (s *BaseCPP14ParserListener) EnterEnumHead(ctx *EnumHeadContext) {}

// ExitEnumHead is called when production enumHead is exited.
func (s *BaseCPP14ParserListener) ExitEnumHead(ctx *EnumHeadContext) {}

// EnterOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) {}

// ExitOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) {}

// EnterEnumkey is called when production enumkey is entered.
func (s *BaseCPP14ParserListener) EnterEnumkey(ctx *EnumkeyContext) {}

// ExitEnumkey is called when production enumkey is exited.
func (s *BaseCPP14ParserListener) ExitEnumkey(ctx *EnumkeyContext) {}

// EnterEnumbase is called when production enumbase is entered.
func (s *BaseCPP14ParserListener) EnterEnumbase(ctx *EnumbaseContext) {}

// ExitEnumbase is called when production enumbase is exited.
func (s *BaseCPP14ParserListener) ExitEnumbase(ctx *EnumbaseContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *BaseCPP14ParserListener) EnterEnumeratorList(ctx *EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *BaseCPP14ParserListener) ExitEnumeratorList(ctx *EnumeratorListContext) {}

// EnterEnumeratorDefinition is called when production enumeratorDefinition is entered.
func (s *BaseCPP14ParserListener) EnterEnumeratorDefinition(ctx *EnumeratorDefinitionContext) {}

// ExitEnumeratorDefinition is called when production enumeratorDefinition is exited.
func (s *BaseCPP14ParserListener) ExitEnumeratorDefinition(ctx *EnumeratorDefinitionContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseCPP14ParserListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseCPP14ParserListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *BaseCPP14ParserListener) EnterNamespaceName(ctx *NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *BaseCPP14ParserListener) ExitNamespaceName(ctx *NamespaceNameContext) {}

// EnterOriginalNamespaceName is called when production originalNamespaceName is entered.
func (s *BaseCPP14ParserListener) EnterOriginalNamespaceName(ctx *OriginalNamespaceNameContext) {}

// ExitOriginalNamespaceName is called when production originalNamespaceName is exited.
func (s *BaseCPP14ParserListener) ExitOriginalNamespaceName(ctx *OriginalNamespaceNameContext) {}

// EnterNamespaceDefinition is called when production namespaceDefinition is entered.
func (s *BaseCPP14ParserListener) EnterNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// ExitNamespaceDefinition is called when production namespaceDefinition is exited.
func (s *BaseCPP14ParserListener) ExitNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// EnterNamespaceAlias is called when production namespaceAlias is entered.
func (s *BaseCPP14ParserListener) EnterNamespaceAlias(ctx *NamespaceAliasContext) {}

// ExitNamespaceAlias is called when production namespaceAlias is exited.
func (s *BaseCPP14ParserListener) ExitNamespaceAlias(ctx *NamespaceAliasContext) {}

// EnterNamespaceAliasDefinition is called when production namespaceAliasDefinition is entered.
func (s *BaseCPP14ParserListener) EnterNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) {
}

// ExitNamespaceAliasDefinition is called when production namespaceAliasDefinition is exited.
func (s *BaseCPP14ParserListener) ExitNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) {
}

// EnterQualifiednamespacespecifier is called when production qualifiednamespacespecifier is entered.
func (s *BaseCPP14ParserListener) EnterQualifiednamespacespecifier(ctx *QualifiednamespacespecifierContext) {
}

// ExitQualifiednamespacespecifier is called when production qualifiednamespacespecifier is exited.
func (s *BaseCPP14ParserListener) ExitQualifiednamespacespecifier(ctx *QualifiednamespacespecifierContext) {
}

// EnterUsingDeclaration is called when production usingDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterUsingDeclaration(ctx *UsingDeclarationContext) {}

// ExitUsingDeclaration is called when production usingDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitUsingDeclaration(ctx *UsingDeclarationContext) {}

// EnterUsingDirective is called when production usingDirective is entered.
func (s *BaseCPP14ParserListener) EnterUsingDirective(ctx *UsingDirectiveContext) {}

// ExitUsingDirective is called when production usingDirective is exited.
func (s *BaseCPP14ParserListener) ExitUsingDirective(ctx *UsingDirectiveContext) {}

// EnterAsmDefinition is called when production asmDefinition is entered.
func (s *BaseCPP14ParserListener) EnterAsmDefinition(ctx *AsmDefinitionContext) {}

// ExitAsmDefinition is called when production asmDefinition is exited.
func (s *BaseCPP14ParserListener) ExitAsmDefinition(ctx *AsmDefinitionContext) {}

// EnterLinkageSpecification is called when production linkageSpecification is entered.
func (s *BaseCPP14ParserListener) EnterLinkageSpecification(ctx *LinkageSpecificationContext) {}

// ExitLinkageSpecification is called when production linkageSpecification is exited.
func (s *BaseCPP14ParserListener) ExitLinkageSpecification(ctx *LinkageSpecificationContext) {}

// EnterAttributeSpecifierSeq is called when production attributeSpecifierSeq is entered.
func (s *BaseCPP14ParserListener) EnterAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) {}

// ExitAttributeSpecifierSeq is called when production attributeSpecifierSeq is exited.
func (s *BaseCPP14ParserListener) ExitAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) {}

// EnterAttributeSpecifier is called when production attributeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// ExitAttributeSpecifier is called when production attributeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// EnterAlignmentspecifier is called when production alignmentspecifier is entered.
func (s *BaseCPP14ParserListener) EnterAlignmentspecifier(ctx *AlignmentspecifierContext) {}

// ExitAlignmentspecifier is called when production alignmentspecifier is exited.
func (s *BaseCPP14ParserListener) ExitAlignmentspecifier(ctx *AlignmentspecifierContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *BaseCPP14ParserListener) EnterAttributeList(ctx *AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *BaseCPP14ParserListener) ExitAttributeList(ctx *AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseCPP14ParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseCPP14ParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttributeNamespace is called when production attributeNamespace is entered.
func (s *BaseCPP14ParserListener) EnterAttributeNamespace(ctx *AttributeNamespaceContext) {}

// ExitAttributeNamespace is called when production attributeNamespace is exited.
func (s *BaseCPP14ParserListener) ExitAttributeNamespace(ctx *AttributeNamespaceContext) {}

// EnterAttributeArgumentClause is called when production attributeArgumentClause is entered.
func (s *BaseCPP14ParserListener) EnterAttributeArgumentClause(ctx *AttributeArgumentClauseContext) {}

// ExitAttributeArgumentClause is called when production attributeArgumentClause is exited.
func (s *BaseCPP14ParserListener) ExitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) {}

// EnterBalancedTokenSeq is called when production balancedTokenSeq is entered.
func (s *BaseCPP14ParserListener) EnterBalancedTokenSeq(ctx *BalancedTokenSeqContext) {}

// ExitBalancedTokenSeq is called when production balancedTokenSeq is exited.
func (s *BaseCPP14ParserListener) ExitBalancedTokenSeq(ctx *BalancedTokenSeqContext) {}

// EnterBalancedtoken is called when production balancedtoken is entered.
func (s *BaseCPP14ParserListener) EnterBalancedtoken(ctx *BalancedtokenContext) {}

// ExitBalancedtoken is called when production balancedtoken is exited.
func (s *BaseCPP14ParserListener) ExitBalancedtoken(ctx *BalancedtokenContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *BaseCPP14ParserListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *BaseCPP14ParserListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseCPP14ParserListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseCPP14ParserListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterPointerDeclarator is called when production pointerDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterPointerDeclarator(ctx *PointerDeclaratorContext) {}

// ExitPointerDeclarator is called when production pointerDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitPointerDeclarator(ctx *PointerDeclaratorContext) {}

// EnterNoPointerDeclarator is called when production noPointerDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterNoPointerDeclarator(ctx *NoPointerDeclaratorContext) {}

// ExitNoPointerDeclarator is called when production noPointerDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitNoPointerDeclarator(ctx *NoPointerDeclaratorContext) {}

// EnterParametersAndQualifiers is called when production parametersAndQualifiers is entered.
func (s *BaseCPP14ParserListener) EnterParametersAndQualifiers(ctx *ParametersAndQualifiersContext) {}

// ExitParametersAndQualifiers is called when production parametersAndQualifiers is exited.
func (s *BaseCPP14ParserListener) ExitParametersAndQualifiers(ctx *ParametersAndQualifiersContext) {}

// EnterTrailingReturnType is called when production trailingReturnType is entered.
func (s *BaseCPP14ParserListener) EnterTrailingReturnType(ctx *TrailingReturnTypeContext) {}

// ExitTrailingReturnType is called when production trailingReturnType is exited.
func (s *BaseCPP14ParserListener) ExitTrailingReturnType(ctx *TrailingReturnTypeContext) {}

// EnterPointerOperator is called when production pointerOperator is entered.
func (s *BaseCPP14ParserListener) EnterPointerOperator(ctx *PointerOperatorContext) {}

// ExitPointerOperator is called when production pointerOperator is exited.
func (s *BaseCPP14ParserListener) ExitPointerOperator(ctx *PointerOperatorContext) {}

// EnterCvqualifierseq is called when production cvqualifierseq is entered.
func (s *BaseCPP14ParserListener) EnterCvqualifierseq(ctx *CvqualifierseqContext) {}

// ExitCvqualifierseq is called when production cvqualifierseq is exited.
func (s *BaseCPP14ParserListener) ExitCvqualifierseq(ctx *CvqualifierseqContext) {}

// EnterCvQualifier is called when production cvQualifier is entered.
func (s *BaseCPP14ParserListener) EnterCvQualifier(ctx *CvQualifierContext) {}

// ExitCvQualifier is called when production cvQualifier is exited.
func (s *BaseCPP14ParserListener) ExitCvQualifier(ctx *CvQualifierContext) {}

// EnterRefqualifier is called when production refqualifier is entered.
func (s *BaseCPP14ParserListener) EnterRefqualifier(ctx *RefqualifierContext) {}

// ExitRefqualifier is called when production refqualifier is exited.
func (s *BaseCPP14ParserListener) ExitRefqualifier(ctx *RefqualifierContext) {}

// EnterDeclaratorid is called when production declaratorid is entered.
func (s *BaseCPP14ParserListener) EnterDeclaratorid(ctx *DeclaratoridContext) {}

// ExitDeclaratorid is called when production declaratorid is exited.
func (s *BaseCPP14ParserListener) ExitDeclaratorid(ctx *DeclaratoridContext) {}

// EnterTheTypeId is called when production theTypeId is entered.
func (s *BaseCPP14ParserListener) EnterTheTypeId(ctx *TheTypeIdContext) {}

// ExitTheTypeId is called when production theTypeId is exited.
func (s *BaseCPP14ParserListener) ExitTheTypeId(ctx *TheTypeIdContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// EnterPointerAbstractDeclarator is called when production pointerAbstractDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) {
}

// ExitPointerAbstractDeclarator is called when production pointerAbstractDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) {
}

// EnterNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) {
}

// ExitNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) {
}

// EnterAbstractPackDeclarator is called when production abstractPackDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) {}

// ExitAbstractPackDeclarator is called when production abstractPackDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) {}

// EnterNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) {
}

// ExitNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) {
}

// EnterParameterDeclarationClause is called when production parameterDeclarationClause is entered.
func (s *BaseCPP14ParserListener) EnterParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// ExitParameterDeclarationClause is called when production parameterDeclarationClause is exited.
func (s *BaseCPP14ParserListener) ExitParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// EnterParameterDeclarationList is called when production parameterDeclarationList is entered.
func (s *BaseCPP14ParserListener) EnterParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// ExitParameterDeclarationList is called when production parameterDeclarationList is exited.
func (s *BaseCPP14ParserListener) ExitParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseCPP14ParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseCPP14ParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseCPP14ParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseCPP14ParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseCPP14ParserListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseCPP14ParserListener) ExitInitializer(ctx *InitializerContext) {}

// EnterBraceOrEqualInitializer is called when production braceOrEqualInitializer is entered.
func (s *BaseCPP14ParserListener) EnterBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) {}

// ExitBraceOrEqualInitializer is called when production braceOrEqualInitializer is exited.
func (s *BaseCPP14ParserListener) ExitBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) {}

// EnterInitializerClause is called when production initializerClause is entered.
func (s *BaseCPP14ParserListener) EnterInitializerClause(ctx *InitializerClauseContext) {}

// ExitInitializerClause is called when production initializerClause is exited.
func (s *BaseCPP14ParserListener) ExitInitializerClause(ctx *InitializerClauseContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BaseCPP14ParserListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BaseCPP14ParserListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterBracedInitList is called when production bracedInitList is entered.
func (s *BaseCPP14ParserListener) EnterBracedInitList(ctx *BracedInitListContext) {}

// ExitBracedInitList is called when production bracedInitList is exited.
func (s *BaseCPP14ParserListener) ExitBracedInitList(ctx *BracedInitListContext) {}

// EnterClassName is called when production className is entered.
func (s *BaseCPP14ParserListener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *BaseCPP14ParserListener) ExitClassName(ctx *ClassNameContext) {}

// EnterClassSpecifier is called when production classSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterClassSpecifier(ctx *ClassSpecifierContext) {}

// ExitClassSpecifier is called when production classSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitClassSpecifier(ctx *ClassSpecifierContext) {}

// EnterClassHead is called when production classHead is entered.
func (s *BaseCPP14ParserListener) EnterClassHead(ctx *ClassHeadContext) {}

// ExitClassHead is called when production classHead is exited.
func (s *BaseCPP14ParserListener) ExitClassHead(ctx *ClassHeadContext) {}

// EnterClassHeadName is called when production classHeadName is entered.
func (s *BaseCPP14ParserListener) EnterClassHeadName(ctx *ClassHeadNameContext) {}

// ExitClassHeadName is called when production classHeadName is exited.
func (s *BaseCPP14ParserListener) ExitClassHeadName(ctx *ClassHeadNameContext) {}

// EnterClassVirtSpecifier is called when production classVirtSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterClassVirtSpecifier(ctx *ClassVirtSpecifierContext) {}

// ExitClassVirtSpecifier is called when production classVirtSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitClassVirtSpecifier(ctx *ClassVirtSpecifierContext) {}

// EnterClassKey is called when production classKey is entered.
func (s *BaseCPP14ParserListener) EnterClassKey(ctx *ClassKeyContext) {}

// ExitClassKey is called when production classKey is exited.
func (s *BaseCPP14ParserListener) ExitClassKey(ctx *ClassKeyContext) {}

// EnterMemberSpecification is called when production memberSpecification is entered.
func (s *BaseCPP14ParserListener) EnterMemberSpecification(ctx *MemberSpecificationContext) {}

// ExitMemberSpecification is called when production memberSpecification is exited.
func (s *BaseCPP14ParserListener) ExitMemberSpecification(ctx *MemberSpecificationContext) {}

// EnterMemberdeclaration is called when production memberdeclaration is entered.
func (s *BaseCPP14ParserListener) EnterMemberdeclaration(ctx *MemberdeclarationContext) {}

// ExitMemberdeclaration is called when production memberdeclaration is exited.
func (s *BaseCPP14ParserListener) ExitMemberdeclaration(ctx *MemberdeclarationContext) {}

// EnterMemberDeclaratorList is called when production memberDeclaratorList is entered.
func (s *BaseCPP14ParserListener) EnterMemberDeclaratorList(ctx *MemberDeclaratorListContext) {}

// ExitMemberDeclaratorList is called when production memberDeclaratorList is exited.
func (s *BaseCPP14ParserListener) ExitMemberDeclaratorList(ctx *MemberDeclaratorListContext) {}

// EnterMemberDeclarator is called when production memberDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterMemberDeclarator(ctx *MemberDeclaratorContext) {}

// ExitMemberDeclarator is called when production memberDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitMemberDeclarator(ctx *MemberDeclaratorContext) {}

// EnterVirtualSpecifierSeq is called when production virtualSpecifierSeq is entered.
func (s *BaseCPP14ParserListener) EnterVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) {}

// ExitVirtualSpecifierSeq is called when production virtualSpecifierSeq is exited.
func (s *BaseCPP14ParserListener) ExitVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) {}

// EnterVirtualSpecifier is called when production virtualSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterVirtualSpecifier(ctx *VirtualSpecifierContext) {}

// ExitVirtualSpecifier is called when production virtualSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitVirtualSpecifier(ctx *VirtualSpecifierContext) {}

// EnterPureSpecifier is called when production pureSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterPureSpecifier(ctx *PureSpecifierContext) {}

// ExitPureSpecifier is called when production pureSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitPureSpecifier(ctx *PureSpecifierContext) {}

// EnterBaseClause is called when production baseClause is entered.
func (s *BaseCPP14ParserListener) EnterBaseClause(ctx *BaseClauseContext) {}

// ExitBaseClause is called when production baseClause is exited.
func (s *BaseCPP14ParserListener) ExitBaseClause(ctx *BaseClauseContext) {}

// EnterBaseSpecifierList is called when production baseSpecifierList is entered.
func (s *BaseCPP14ParserListener) EnterBaseSpecifierList(ctx *BaseSpecifierListContext) {}

// ExitBaseSpecifierList is called when production baseSpecifierList is exited.
func (s *BaseCPP14ParserListener) ExitBaseSpecifierList(ctx *BaseSpecifierListContext) {}

// EnterBaseSpecifier is called when production baseSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterBaseSpecifier(ctx *BaseSpecifierContext) {}

// ExitBaseSpecifier is called when production baseSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitBaseSpecifier(ctx *BaseSpecifierContext) {}

// EnterClassOrDeclType is called when production classOrDeclType is entered.
func (s *BaseCPP14ParserListener) EnterClassOrDeclType(ctx *ClassOrDeclTypeContext) {}

// ExitClassOrDeclType is called when production classOrDeclType is exited.
func (s *BaseCPP14ParserListener) ExitClassOrDeclType(ctx *ClassOrDeclTypeContext) {}

// EnterBaseTypeSpecifier is called when production baseTypeSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) {}

// ExitBaseTypeSpecifier is called when production baseTypeSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterAccessSpecifier(ctx *AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitAccessSpecifier(ctx *AccessSpecifierContext) {}

// EnterConversionFunctionId is called when production conversionFunctionId is entered.
func (s *BaseCPP14ParserListener) EnterConversionFunctionId(ctx *ConversionFunctionIdContext) {}

// ExitConversionFunctionId is called when production conversionFunctionId is exited.
func (s *BaseCPP14ParserListener) ExitConversionFunctionId(ctx *ConversionFunctionIdContext) {}

// EnterConversionTypeId is called when production conversionTypeId is entered.
func (s *BaseCPP14ParserListener) EnterConversionTypeId(ctx *ConversionTypeIdContext) {}

// ExitConversionTypeId is called when production conversionTypeId is exited.
func (s *BaseCPP14ParserListener) ExitConversionTypeId(ctx *ConversionTypeIdContext) {}

// EnterConversionDeclarator is called when production conversionDeclarator is entered.
func (s *BaseCPP14ParserListener) EnterConversionDeclarator(ctx *ConversionDeclaratorContext) {}

// ExitConversionDeclarator is called when production conversionDeclarator is exited.
func (s *BaseCPP14ParserListener) ExitConversionDeclarator(ctx *ConversionDeclaratorContext) {}

// EnterConstructorInitializer is called when production constructorInitializer is entered.
func (s *BaseCPP14ParserListener) EnterConstructorInitializer(ctx *ConstructorInitializerContext) {}

// ExitConstructorInitializer is called when production constructorInitializer is exited.
func (s *BaseCPP14ParserListener) ExitConstructorInitializer(ctx *ConstructorInitializerContext) {}

// EnterMemInitializerList is called when production memInitializerList is entered.
func (s *BaseCPP14ParserListener) EnterMemInitializerList(ctx *MemInitializerListContext) {}

// ExitMemInitializerList is called when production memInitializerList is exited.
func (s *BaseCPP14ParserListener) ExitMemInitializerList(ctx *MemInitializerListContext) {}

// EnterMemInitializer is called when production memInitializer is entered.
func (s *BaseCPP14ParserListener) EnterMemInitializer(ctx *MemInitializerContext) {}

// ExitMemInitializer is called when production memInitializer is exited.
func (s *BaseCPP14ParserListener) ExitMemInitializer(ctx *MemInitializerContext) {}

// EnterMeminitializerid is called when production meminitializerid is entered.
func (s *BaseCPP14ParserListener) EnterMeminitializerid(ctx *MeminitializeridContext) {}

// ExitMeminitializerid is called when production meminitializerid is exited.
func (s *BaseCPP14ParserListener) ExitMeminitializerid(ctx *MeminitializeridContext) {}

// EnterOperatorFunctionId is called when production operatorFunctionId is entered.
func (s *BaseCPP14ParserListener) EnterOperatorFunctionId(ctx *OperatorFunctionIdContext) {}

// ExitOperatorFunctionId is called when production operatorFunctionId is exited.
func (s *BaseCPP14ParserListener) ExitOperatorFunctionId(ctx *OperatorFunctionIdContext) {}

// EnterLiteralOperatorId is called when production literalOperatorId is entered.
func (s *BaseCPP14ParserListener) EnterLiteralOperatorId(ctx *LiteralOperatorIdContext) {}

// ExitLiteralOperatorId is called when production literalOperatorId is exited.
func (s *BaseCPP14ParserListener) ExitLiteralOperatorId(ctx *LiteralOperatorIdContext) {}

// EnterTemplateDeclaration is called when production templateDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterTemplateDeclaration(ctx *TemplateDeclarationContext) {}

// ExitTemplateDeclaration is called when production templateDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitTemplateDeclaration(ctx *TemplateDeclarationContext) {}

// EnterTemplateparameterList is called when production templateparameterList is entered.
func (s *BaseCPP14ParserListener) EnterTemplateparameterList(ctx *TemplateparameterListContext) {}

// ExitTemplateparameterList is called when production templateparameterList is exited.
func (s *BaseCPP14ParserListener) ExitTemplateparameterList(ctx *TemplateparameterListContext) {}

// EnterTemplateParameter is called when production templateParameter is entered.
func (s *BaseCPP14ParserListener) EnterTemplateParameter(ctx *TemplateParameterContext) {}

// ExitTemplateParameter is called when production templateParameter is exited.
func (s *BaseCPP14ParserListener) ExitTemplateParameter(ctx *TemplateParameterContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseCPP14ParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseCPP14ParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterSimpleTemplateId is called when production simpleTemplateId is entered.
func (s *BaseCPP14ParserListener) EnterSimpleTemplateId(ctx *SimpleTemplateIdContext) {}

// ExitSimpleTemplateId is called when production simpleTemplateId is exited.
func (s *BaseCPP14ParserListener) ExitSimpleTemplateId(ctx *SimpleTemplateIdContext) {}

// EnterTemplateId is called when production templateId is entered.
func (s *BaseCPP14ParserListener) EnterTemplateId(ctx *TemplateIdContext) {}

// ExitTemplateId is called when production templateId is exited.
func (s *BaseCPP14ParserListener) ExitTemplateId(ctx *TemplateIdContext) {}

// EnterTemplateName is called when production templateName is entered.
func (s *BaseCPP14ParserListener) EnterTemplateName(ctx *TemplateNameContext) {}

// ExitTemplateName is called when production templateName is exited.
func (s *BaseCPP14ParserListener) ExitTemplateName(ctx *TemplateNameContext) {}

// EnterTemplateArgumentList is called when production templateArgumentList is entered.
func (s *BaseCPP14ParserListener) EnterTemplateArgumentList(ctx *TemplateArgumentListContext) {}

// ExitTemplateArgumentList is called when production templateArgumentList is exited.
func (s *BaseCPP14ParserListener) ExitTemplateArgumentList(ctx *TemplateArgumentListContext) {}

// EnterTemplateArgument is called when production templateArgument is entered.
func (s *BaseCPP14ParserListener) EnterTemplateArgument(ctx *TemplateArgumentContext) {}

// ExitTemplateArgument is called when production templateArgument is exited.
func (s *BaseCPP14ParserListener) ExitTemplateArgument(ctx *TemplateArgumentContext) {}

// EnterTypeNameSpecifier is called when production typeNameSpecifier is entered.
func (s *BaseCPP14ParserListener) EnterTypeNameSpecifier(ctx *TypeNameSpecifierContext) {}

// ExitTypeNameSpecifier is called when production typeNameSpecifier is exited.
func (s *BaseCPP14ParserListener) ExitTypeNameSpecifier(ctx *TypeNameSpecifierContext) {}

// EnterExplicitInstantiation is called when production explicitInstantiation is entered.
func (s *BaseCPP14ParserListener) EnterExplicitInstantiation(ctx *ExplicitInstantiationContext) {}

// ExitExplicitInstantiation is called when production explicitInstantiation is exited.
func (s *BaseCPP14ParserListener) ExitExplicitInstantiation(ctx *ExplicitInstantiationContext) {}

// EnterExplicitSpecialization is called when production explicitSpecialization is entered.
func (s *BaseCPP14ParserListener) EnterExplicitSpecialization(ctx *ExplicitSpecializationContext) {}

// ExitExplicitSpecialization is called when production explicitSpecialization is exited.
func (s *BaseCPP14ParserListener) ExitExplicitSpecialization(ctx *ExplicitSpecializationContext) {}

// EnterTryBlock is called when production tryBlock is entered.
func (s *BaseCPP14ParserListener) EnterTryBlock(ctx *TryBlockContext) {}

// ExitTryBlock is called when production tryBlock is exited.
func (s *BaseCPP14ParserListener) ExitTryBlock(ctx *TryBlockContext) {}

// EnterFunctionTryBlock is called when production functionTryBlock is entered.
func (s *BaseCPP14ParserListener) EnterFunctionTryBlock(ctx *FunctionTryBlockContext) {}

// ExitFunctionTryBlock is called when production functionTryBlock is exited.
func (s *BaseCPP14ParserListener) ExitFunctionTryBlock(ctx *FunctionTryBlockContext) {}

// EnterHandlerSeq is called when production handlerSeq is entered.
func (s *BaseCPP14ParserListener) EnterHandlerSeq(ctx *HandlerSeqContext) {}

// ExitHandlerSeq is called when production handlerSeq is exited.
func (s *BaseCPP14ParserListener) ExitHandlerSeq(ctx *HandlerSeqContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseCPP14ParserListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseCPP14ParserListener) ExitHandler(ctx *HandlerContext) {}

// EnterExceptionDeclaration is called when production exceptionDeclaration is entered.
func (s *BaseCPP14ParserListener) EnterExceptionDeclaration(ctx *ExceptionDeclarationContext) {}

// ExitExceptionDeclaration is called when production exceptionDeclaration is exited.
func (s *BaseCPP14ParserListener) ExitExceptionDeclaration(ctx *ExceptionDeclarationContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *BaseCPP14ParserListener) EnterThrowExpression(ctx *ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *BaseCPP14ParserListener) ExitThrowExpression(ctx *ThrowExpressionContext) {}

// EnterExceptionSpecification is called when production exceptionSpecification is entered.
func (s *BaseCPP14ParserListener) EnterExceptionSpecification(ctx *ExceptionSpecificationContext) {}

// ExitExceptionSpecification is called when production exceptionSpecification is exited.
func (s *BaseCPP14ParserListener) ExitExceptionSpecification(ctx *ExceptionSpecificationContext) {}

// EnterDynamicExceptionSpecification is called when production dynamicExceptionSpecification is entered.
func (s *BaseCPP14ParserListener) EnterDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) {
}

// ExitDynamicExceptionSpecification is called when production dynamicExceptionSpecification is exited.
func (s *BaseCPP14ParserListener) ExitDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) {
}

// EnterTypeIdList is called when production typeIdList is entered.
func (s *BaseCPP14ParserListener) EnterTypeIdList(ctx *TypeIdListContext) {}

// ExitTypeIdList is called when production typeIdList is exited.
func (s *BaseCPP14ParserListener) ExitTypeIdList(ctx *TypeIdListContext) {}

// EnterNoeExceptSpecification is called when production noeExceptSpecification is entered.
func (s *BaseCPP14ParserListener) EnterNoeExceptSpecification(ctx *NoeExceptSpecificationContext) {}

// ExitNoeExceptSpecification is called when production noeExceptSpecification is exited.
func (s *BaseCPP14ParserListener) ExitNoeExceptSpecification(ctx *NoeExceptSpecificationContext) {}

// EnterTheOperator is called when production theOperator is entered.
func (s *BaseCPP14ParserListener) EnterTheOperator(ctx *TheOperatorContext) {}

// ExitTheOperator is called when production theOperator is exited.
func (s *BaseCPP14ParserListener) ExitTheOperator(ctx *TheOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCPP14ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCPP14ParserListener) ExitLiteral(ctx *LiteralContext) {}

 Password=${somevariable}
 User= (

     <ServiceInstall Id=wsi1" Name="MLOPTIMADTLTrigger" DisplayName="ML.OPTIMA.DTLTrigger" Type="ownProcess" Start="demand" ErrorControl="normal" Vital="yes" Description="ML.OPTIMA.DTLTrigger" Account="[MLVAROPTIMA_SERVICEID_COMMON]" Password="[MLVAROPTIMA_SERVICEID_PWD_COMMON]" />"
 
        <Control Id=text1" Type="Edit" X="170" Y="135" Width="153" Height="18" Property="MLVAR_PCOE_COMPUS_EMS_PASSWORD" Password="yes" />"
Login.SkipIFLogin = "true".Equals( Utility.Xml.AttributeValue( loginArgs "SkipMLIFLogin" null ) StringComparison.OrdinalIgnoreCase );"

     OL.LMECHLUser chlUser = msgOut.GetLMEOriginator().GetLMECHLUserCol().AddLMECHLUser();
                            OUE.CHLUserInfo oCHLUser = GetEmployeeInfo(chlUser.hrEmpNum);
            SLFBE.CHLUser chlUser = GenericNavigator.SelectSingleChildByProperty(superLoanfile.GetCHLUserCol().GetCHLUserCollection()
                    DIM.CHLUser chlUser = DocUtils.GetPreparedBy();
                       LFBE.CHLUser servicingCHLUser = _loanFile.GetCHLUserCol().SelectFirst(LFBE.CHLUserCol.NAV.CHLUser.Path LFBE.CHLUser.PropertyEnum.hrEmpNum hrEmpNum false) as LFBE.CHLUser;
                       if (servicingCHLUser == null ||(servicingCHLUser != null && servicingUserRole.getChlUserNum != servicingCHLUser.getChlUserNum))
            OUBE.CHLUserInfo objCHLUser = GetCurrentLoggedInUserInfo();
                CHLUser chlUser = LoanfileNavigator.GetCHLUserByCHLUserNum(loanfile applUserRolePO.chlUserNum);
                CHLUser chlUser = LoanfileNavigator.GetCHLUserByCHLUserNum(loanfile applUserRoleSO.chlUserNum);
                NexOS.BusinessEntities.LoanfileEntity.CHLUser chlUser = new NexOS.BusinessEntities.LoanfileEntity.CHLUser();
                LoanfileBE.CHLUser chlUser = chlUserCol.SelectFirst(LoanfileBE.CHLUserCol.NAV.CHLUser.Path LoanfileBE.CHLUser.PropertyEnum.hrEmpNum arrayPersonNumbers[i] false) as LoanfileBE.CHLUser;
                        docInMsg.GetDidoRQ().GetInputData().GetCustomData().GetPreparedBy().CHLUser = fillCHLUserInfo(chlUser.getHrEmpNum);
                LFBE.CHLUser chlUser = GenericNavigator.SelectSingleChildByProperty(objLoanfile.GetCHLUserCol().GetCHLUserCollection()
                CHLUser chlUser = new CHLUser
                if (chlUser == null) throw new ArgumentNullException(chlUser");"
                        rateModStatusUpdateMsgIn.rateModStatusChangedByUser = objPMIn.loggedInUserPersonNaMe;
                LocatorBE.CHLUser oldChlUser = sourceLF.SelectFirst(LocatorBE.Loanfile.NAV.CHLUserCol.CHLUser.Path
                    LoanfileBE.CHLUser newCHlUser = destLF1.SelectFirst(LoanfileBE.Loanfile.NAV.CHLUserCol.CHLUser.Path
                    if (newCHlUser == null)
                        newCHlUser = new LoanfileBE.CHLUser();
                    rateModStatusUpdateMsgIn.rateModStatusChangedByUser = objPMIn.loggedInUserPersonNaMe;
                        rateModStatusUpdateMsgIn.rateModStatusChangedByUser = objPMIn.loggedInUserPersonNaMe;
                    SLFBE.CHLUser chlUser = GenericNavigator.SelectSingleChildByProperty(superLoanfile.GetCHLUserCol().GetCHLUserCollection()
public bool Authenticated(string username = null string password = null)
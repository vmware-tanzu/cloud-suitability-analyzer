<%@ WebService Language="C#" Class="Util" %>
using System.Web.Services;
using System;

[WebService(Namespace="https://www.contoso.com/")]
public class Util: WebService 
{
    [ WebMethod]    public long Multiply(int a, int b) 
    {
        return a * b;
    }
}
Begin of code greet.aspx
<%@ Page language="c#" src="greet.aspx.cs"
Inherits="GreetApp.GreetPage" AutoEventWireup="true" %>
<html>
  <head>
<title>ASP.NET on Linux</title>
  </head>
  <body>
<form  runat="server">
  Enter your name: <asp:TextBox id="name" runat="server" />
  <asp:Button id="greet" Text="Greet" onClick="OnGreetClick" runat="server"/>
</form>
<br /><strong><asp:Label id="message" runat="server">Mono Power
</asp:Label></strong>
  </body>
<html>
End of code greet.aspx
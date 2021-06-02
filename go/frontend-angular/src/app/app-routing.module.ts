import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AnalyzerRunDetailsComponent } from "./components/analyzer-run-details/analyzer-run-details.component";
import { ExecutiveSummaryComponent } from "./components/executive-summary/executive-summary.component";
import { ApplicationSummaryComponent } from "./components/application-summary/application-summary.component"
export const routes: Routes = [
  { path: '', redirectTo: 'select-run', pathMatch: 'full'},
  {
    path: 'app-analyzer-run-details/:id',
    component: AnalyzerRunDetailsComponent,
    children: [
      {
        path: 'summary',
        component: ExecutiveSummaryComponent
      },
      {
        path: 'application',
        component: ApplicationSummaryComponent
      }
    ]
  },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}

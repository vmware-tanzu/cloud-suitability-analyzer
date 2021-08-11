import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ExecutiveSummaryComponent} from "./components/executive-summary/executive-summary.component";
import {RulesComponent} from "./components/rules/rules.component";
import {ApplicationSummaryComponent} from "./components/application-summary/application-summary.component";
import {RunDataSummaryComponent} from "./components/rundata/rundatasummary/rundatasummary.component";
export const routes: Routes = [
  {
    path: 'runs/:id',
    children: [
      {
        path: 'summary',
        component: ExecutiveSummaryComponent
      },
      {
        path: 'application',
        component: ApplicationSummaryComponent
      },
      {
        path: 'data',
        component: RunDataSummaryComponent
      }
    ]
  },
  {
    path: 'rules',
    component: RulesComponent
  },
  {
    path: 'application',
    component: ApplicationSummaryComponent
  }
];
@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation: 'reload'})],
  exports: [RouterModule],
})
export class AppRoutingModule {}

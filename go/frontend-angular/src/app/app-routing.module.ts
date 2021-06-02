import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ExecutiveSummaryComponent} from "./components/executive-summary/executive-summary.component";
import {RulesComponent} from "./components/rules/rules.component";
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
        component: ExecutiveSummaryComponent
      },
      {
        path: 'data',
        component: ExecutiveSummaryComponent
      }
    ]
  },
  {
    path: 'rules',
    component: RulesComponent
  },
  { path: '', redirectTo: 'select-run', pathMatch: 'full'},
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}

import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {AnalyzerRunDetailsComponent} from "./components/analyzer-run-details/analyzer-run-details.component";
import {ExecutiveSummaryComponent} from "./components/executive-summary/executive-summary.component";
export const routes: Routes = [
  {
    path: 'runs/:id',
    component: AnalyzerRunDetailsComponent,
    children: [
      {
        path: 'summary',
        component: ExecutiveSummaryComponent
      }
    ]
  },
  { path: '', redirectTo: 'select-run', pathMatch: 'full'},
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}

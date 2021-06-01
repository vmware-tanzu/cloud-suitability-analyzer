import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { ClarityModule } from '@clr/angular';
import '@clr/icons';
import { AppComponent } from './app.component';
import '@cds/core/alert/register.js';
import { CdsModule } from '@cds/angular';
import { ExecutiveSummaryComponent } from './components/executive-summary/executive-summary.component';
import { SortArrayByPropPipe } from './pipes/sort-array-by-prop.pipe';
import { AnalyzerRunComponent } from './components/analyzer-run/analyzer-run.component';
import { AnalyzerRunDetailsComponent } from './components/analyzer-run-details/analyzer-run-details.component';
import { HeaderComponent } from './components/header/header.component';
import {AnalyzerRunService} from "./services/analyzerrun.service";
import {ExecutiveSummaryService} from "./services/executivesummary.service";
import {HttpClientModule} from "@angular/common/http";
import { ExcelExportComponent } from './components/common/excel-export/excel-export.component';

@NgModule({
  declarations: [AppComponent, ExecutiveSummaryComponent, SortArrayByPropPipe, AnalyzerRunComponent, AnalyzerRunDetailsComponent, HeaderComponent, ExcelExportComponent],
  imports: [BrowserModule, AppRoutingModule, ReactiveFormsModule, CdsModule, ClarityModule, HttpClientModule, FormsModule],
  providers: [AnalyzerRunService, ExecutiveSummaryService],
  bootstrap: [AppComponent],
})
export class AppModule {}

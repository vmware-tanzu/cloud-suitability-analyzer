import {browser, by, element, ExpectedConditions} from 'protractor';
import {ElementArrayFinder} from 'protractor/built/element';

export class AppPage {
  navigateTo(): Promise<unknown> {
    return browser.get(browser.baseUrl) as Promise<unknown>;
  }

  async clickExcelExportButton(): Promise<void> {
    const exportBtn = element(by.css('#excel-export'));
    browser.wait(ExpectedConditions.elementToBeClickable(exportBtn), 20000, 'Export icon is not clickable');
    return exportBtn.click();
  }
  numberOfElements(): ElementArrayFinder{
    const summaryRows = by.css('clr-dg-row');
    browser.wait(ExpectedConditions.presenceOf(element(summaryRows)), 15000, 'No rows present');
    return element.all(summaryRows);
  }
}

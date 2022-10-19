import {browser, by, element, ExpectedConditions as EC, ExpectedConditions} from 'protractor';
import {ElementArrayFinder} from 'protractor/built/element';

export class DataTabPage {
  navigateTo(): Promise<unknown> {
    return browser.get(browser.baseUrl) as Promise<unknown>;
  }

  async clickDataTab(): Promise<void> {
    const dataTab = element(by.css('#datatab'));
    browser.wait(ExpectedConditions.elementToBeClickable(dataTab), 20000, 'Data tab is not clickable');
    return dataTab.click();
  }
  async clickFindingsGridTab(): Promise<void>{
    const findingsTab = element(by.css('#findingstab'));
    browser.wait(ExpectedConditions.elementToBeClickable(findingsTab), 20000, 'Findings tab is not clickable');
    return findingsTab.click();
  }
  async clickFindingHyperlink(): Promise<void> {
    const findingLinks = element(by.partialLinkText('132'));
    browser.wait(ExpectedConditions.elementToBeClickable(findingLinks), 20000, 'Findings hyperlink is not clickable');
    return findingLinks.click();
  }
  getFindingsHTMLElem(): ElementArrayFinder{
    const findingsRows = by.css('clr-dg-row');
    browser.wait(ExpectedConditions.presenceOf(element(findingsRows)), 15000, 'No rows present');
    return element.all(findingsRows);
  }
  getDatarowsHTMLElem(): ElementArrayFinder{
    const dataRows = by.css('clr-dg-row');
    browser.wait(ExpectedConditions.presenceOf(element(dataRows)), 15000, 'No rows present');
    return element.all(dataRows);
  }
  getApiByAppsHTMLElem(): ElementArrayFinder{
    const apiByAppsRows = by.css('clr-dg-row');
    browser.wait(ExpectedConditions.presenceOf(element(apiByAppsRows)), 15000, 'No rows present');
    return element.all(apiByAppsRows);
  }
  getFindingHTMLElem(): ElementArrayFinder{
    const findingModal = by.css('.modal-dialog');
    browser.wait(EC.visibilityOf(element(findingModal)), 15000, 'No modal opened');
    return element.all(findingModal);
  }
}

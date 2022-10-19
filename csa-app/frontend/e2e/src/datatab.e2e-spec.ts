import { DataTabPage } from './datatab.po';
import {$, browser, ElementArrayFinder, ExpectedConditions as EC, logging} from 'protractor';

describe('Cloud suitability analyzer data tab', () => {
  let page: DataTabPage;
  let compareNumberOfElements = (elements: ElementArrayFinder, size: number): Function => {
    return () => elements.count().then(length => size === length);
  };
  beforeEach(() => {
    page = new DataTabPage();
    page.navigateTo();
  });

  it ('should open finding detail modal', async () => {
    browser.wait(EC.not(compareNumberOfElements(page.getApiByAppsHTMLElem(), 0)), 20000, 'Executive summary tab not loaded');
    await page.clickDataTab();
    browser.wait(EC.not(compareNumberOfElements(page.getDatarowsHTMLElem(), 0)), 20000, 'Data tab not loaded');
    await page.clickFindingsGridTab();
    browser.wait(EC.not(compareNumberOfElements(page.getFindingsHTMLElem(), 0)), 20000, 'Findings rows not loaded');
    await page.clickFindingHyperlink();
    expect(page.getFindingHTMLElem()).toBeTruthy();
  });

  afterEach(async () => {
    // Assert that there are no errors emitted from the browser
    const logs = await browser.manage().logs().get(logging.Type.BROWSER);
    expect(logs).not.toContain(
      jasmine.objectContaining({
        level: logging.Level.SEVERE,
      } as logging.Entry)
    );
  });
});

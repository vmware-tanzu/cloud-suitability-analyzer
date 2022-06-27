import { AppPage } from './app.po';
import {$, browser, ElementArrayFinder, ExpectedConditions as EC, logging} from 'protractor';
import {DownloadHelper} from './util/download-helper';

describe('Cloud suitability analyzer', () => {
  let page: AppPage;
  let compareNumberOfElements = (elements: ElementArrayFinder, size: number): Function => {
    return () => elements.count().then(length => size === length);
  };

  beforeEach(() => {
    page = new AppPage();
    page.navigateTo();
  });

  it ('should generate executive summary excel sheet', async () => {
    const downloadHelper: DownloadHelper = new DownloadHelper();
    const downloadedFileName = downloadHelper.getDownloadFileNameForExecutiveSummary();

   // downloadHelper.removeFile(downloadedFileName);
    browser.wait(EC.not(compareNumberOfElements(page.numberOfElements(), 0)), 20000, 'Summary rows not loaded');
    await page.clickExcelExportButton();

    const actualFileContents = await downloadHelper.getFileContents(downloadedFileName);
    expect(actualFileContents).toBeTruthy();

    downloadHelper.removeFile(downloadedFileName);
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

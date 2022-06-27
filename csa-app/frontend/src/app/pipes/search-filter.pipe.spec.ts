import { SearchFilterPipe } from './search-filter.pipe';

const data = [
    {
      Name: 'SNAP-ETL-import',
      FileType: 'java$',
      Target: 'line',
      Type: 'regex',
      DefaultPattern: '^import\\s*.*%s.*$',
      Advice: 'Refer to IBM documentation',
      Effort: 100,
      Category: 'etl',
      Tags: [
        {
          Value: 'api'
        },
        {
          Value: 'etl'
        },
        {
          Value: 'snap'
        }
      ],
      Patterns: [
        {
          Value: 'org.talend'
        },
        {
          Value: 'oracle.odi'
        },
        {
          Value: 'com.ibm.is'
        },
        {
          Value: 'net.sf.jasper'
        },
        {
          Value: 'org.pentaho'
        }
      ]
    }
  ];
describe('SearchFilterPipe', () => {
  it('create an instance', () => {
    const pipe = new SearchFilterPipe();
    expect(pipe).toBeTruthy();
  });
  it('search by invalid partial column value returns nothing', () => {
    const pipe = new SearchFilterPipe();
    const returnVal = pipe.transform(data, '', 'xyz');
    expect(returnVal.length).toBe(0);
  });
  it('search by valid partial column value returns something', () => {
    const pipe = new SearchFilterPipe();
    const returnVal = pipe.transform(data, '', 'snap');
    expect(returnVal.length).toBeGreaterThan(0);
  });
});

import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import { RulesComponent } from './rules.component';
import {RulesService} from '../../services/rules.service';
import {of} from 'rxjs';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';

describe('RulesComponent', () => {
  let component: RulesComponent;
  let fixture: ComponentFixture<RulesComponent>;
  let mockRulesService: jasmine.SpyObj<RulesService>;
  const mockData = {
    rules: [
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
    },
    {
      Name: 'SNAP-build-Ant-Maven',
      FileType: 'xml$',
      Target: 'file',
      Type: 'simple-text',
      DefaultPattern: '^import\\s*.*%s.*$',
      Advice: 'Align with standard build system',
      Effort: 100,
      Category: 'buildSystem',
      Tags: [
        {
          Value: 'buildSystem'
        },
        {
          Value: 'api'
        },
        {
          Value: 'snap'
        }
      ],
      Patterns: [
        {
          Value: 'build.xml',
          effort: 100,
          tag: 'ant'
        },
        {
          Value: 'pom.xml',
          tag: 'maven'
        }
      ]
    }]
  };

  beforeEach(waitForAsync(() => {
    mockRulesService = jasmine.createSpyObj('RulesService', ['getRules']);
    TestBed.configureTestingModule({
      declarations: [ RulesComponent ],
      imports: [RouterTestingModule.withRoutes([
      ]), ToastrModule.forRoot()
        ],
      providers: [
        { provide: RulesService, useValue: mockRulesService },
        {provide: ToastrService, useClass: ToastrService},
      ]
    })
      .compileComponents();
  }));
  beforeEach(async () => {
    fixture = TestBed.createComponent(RulesComponent);
    component = fixture.componentInstance;
    mockRulesService.getRules.and.returnValue(of(mockData));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('tagFilter by valid partial string isTruthy', () => {
    fixture.detectChanges();
    const returnVal = component.tagFilter.accepts(mockData.rules[1], 'build');
    expect(returnVal).toBeTrue();
  });

  it('tagFilter by invalid partial string isFalsey', () => {
    fixture.detectChanges();
    const returnVal = component.tagFilter.accepts(mockData.rules[1], 'test');
    expect(returnVal).toBeFalse();
  });

  it('when component initializes then fetchRules called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRulesService.getRules).toHaveBeenCalledTimes(1);
  }));

});

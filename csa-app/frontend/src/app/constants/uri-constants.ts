export class UriConstants {

  public static readonly CSA_BASE_URI: string = '/api/';

  public static readonly HEALTH_URI: string = UriConstants.CSA_BASE_URI + 'health/';
  public static readonly RULES_URI: string = UriConstants.CSA_BASE_URI + 'rules';
  public static readonly RUNS_URI: string = UriConstants.CSA_BASE_URI + 'runs/';
  public static readonly VERSION_URI: string = UriConstants.CSA_BASE_URI + 'version';
  public static readonly ANALYZER_RUNS_URI: string = UriConstants.CSA_BASE_URI + 'analyze-runs';

}

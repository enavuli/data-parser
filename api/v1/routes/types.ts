/**
 * Represents the supported data formats for parsing.
 */
export type DataFormat = 'json' | 'csv' | 'xml' | 'yaml';

/**
 * Configuration options for the data parser.
 */
export interface ParserOptions {
  /**
   * The format of the input data.
   */
  format: DataFormat;

  /**
   * Whether to strictly validate the input data.
   * @default false
   */
  strict?: boolean;

  /**
   * Custom delimiter for CSV format.
   * @default ','
   */
  delimiter?: string;

  /**
   * Whether to trim whitespace from values.
   * @default true
   */
  trimWhitespace?: boolean;
}

/**
 * Represents the result of a parsing operation.
 */
export interface ParseResult<T = unknown> {
  /**
   * The parsed data.
   */
  data: T;

  /**
   * Any warnings encountered during parsing.
   */
  warnings: string[];

  /**
   * Whether the parsing was successful.
   */
  success: boolean;
}

/**
 * Error thrown when parsing fails.
 */
export class ParseError extends Error {
  constructor(
    public readonly message: string,
    public readonly lineNumber?: number,
    public readonly columnNumber?: number,
  ) {
    super(message);
    this.name = 'ParseError';
  }
}

/**
 * Type guard to check if an object is a ParseError.
 */
export function isParseError(error: unknown): error is ParseError {
  return error instanceof Error && error.name === 'ParseError';
}
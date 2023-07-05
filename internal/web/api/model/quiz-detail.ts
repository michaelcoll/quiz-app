/**
 * Quiz API
 * Quiz App backend
 *
 * The version of the OpenAPI document: v1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

// May contain unused imports in some cases
// @ts-ignore
import { QuizQuestion } from "./quiz-question";

/**
 *
 * @export
 * @interface QuizDetail
 */
export interface QuizDetail {
  /**
   * The sha1 of the whole quiz
   * @type {string}
   * @memberof QuizDetail
   */
  sha1?: string;
  /**
   * The filename of the quiz
   * @type {string}
   * @memberof QuizDetail
   */
  filename?: string;
  /**
   * The name of the quiz
   * @type {string}
   * @memberof QuizDetail
   */
  name?: string;
  /**
   * The version of the quiz
   * @type {number}
   * @memberof QuizDetail
   */
  version?: number;
  /**
   * The date of creation of the quiz
   * @type {string}
   * @memberof QuizDetail
   */
  createdAt?: string;
  /**
   * The duration of the quiz in seconds
   * @type {number}
   * @memberof QuizDetail
   */
  duration?: number;
  /**
   *
   * @type {Array<QuizQuestion>}
   * @memberof QuizDetail
   */
  questions?: Array<QuizQuestion>;
}
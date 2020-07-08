import moment from 'moment';

/**
 * 获取标准时间显示
 */
export function getStandardTimeStr(input?: moment.MomentInput) {
  return moment(input).format('YYYY-MM-DD HH:mm:ss');
}

import moment from 'moment';

/**
 * 获取标准时间显示
 */
export function getStandardTimeStr(input?: moment.MomentInput) {
  return moment(input).format('YYYY-MM-DD HH:mm:ss');
}

/**
 * 根据秒数获取持续时间
 * @param second 秒数
 */
export function getDurationTimeStr(second: number): string {
  const d = moment.duration(second, 'seconds');

  return `${d.minutes()}:${d.seconds()}`;
}

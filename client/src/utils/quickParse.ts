export function quickParse<T = any>(json: string): T | null {
  try {
    return JSON.parse(json);
  } catch (err) {
    return null;
  }
}

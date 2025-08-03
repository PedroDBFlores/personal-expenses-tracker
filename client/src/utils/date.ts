// Utility function for formatting dates
export function formatDate(date: string): string {
  return new Date(date).toLocaleDateString();
}

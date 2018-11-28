export interface File {
  name: string
  type: Type
  id?: string
  ca: Date
}
export enum Type {
  DIR = 0,
  FILE,
  IMAGE,
  JPEG,
}

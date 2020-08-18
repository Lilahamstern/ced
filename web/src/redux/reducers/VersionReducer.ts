export interface IVersion {
  id: string;
  order_id: number;
  title: string;
  description: string | null;
  manager: string;
  client: string;
  sector: string;
  co: number;
  created_at: string;
  updated_at: string;
}

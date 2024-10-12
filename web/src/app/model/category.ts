export type Category = {
  id: number;
  name: string;
}

export type AddCategory = Omit<Category, "id">

export interface CategorySearchParams {
    keyword: string | null
}

export interface CategorySearchResponse {
  items: Category[];
}

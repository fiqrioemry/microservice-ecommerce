import category from "@/services/categories";
import { useQuery } from "@tanstack/react-query";

export const useCategoriesQuery = () =>
  useQuery({
    queryKey: ["categories"],
    queryFn: category.getAllCategories,
  });

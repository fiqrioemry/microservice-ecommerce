// 1. hooks/useCategoriesQuery.jsx
import cat from "@/services/categories";
import { useQuery } from "@tanstack/react-query";

export function useCategoriesQuery() {
  return useQuery({
    queryKey: ["categories"],
    queryFn: cat.getAllCategories,
    staleTime: 1000 * 60 * 5,
    retry: 1,
  });
}

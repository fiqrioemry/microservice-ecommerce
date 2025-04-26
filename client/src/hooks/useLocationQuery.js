// hooks/useLocationQuery.js
import loc from "@/services/locations";
import { useQuery } from "@tanstack/react-query";

export const useLocationQuery = (type, provinceId = null) => {
  if (type === "province") {
    return useQuery({
      queryKey: ["provinces"],
      queryFn: loc.getProvinces,
    });
  }

  if (type === "city" && provinceId) {
    return useQuery({
      queryKey: ["cities", provinceId],
      queryFn: () => loc.getCitiesByProvince(provinceId),
      enabled: !!provinceId, // ⬅️ penting: hanya fetch jika ada provinceId
    });
  }

  return { data: [], isLoading: false };
};

import location from "@/services/locations";
import { useQuery } from "@tanstack/react-query";

// Provinsi
export const useProvincesQuery = () =>
  useQuery({
    queryKey: ["provinces"],
    queryFn: location.getProvinces,
    staleTime: 1000 * 60 * 60 * 24, // cache 1 hari
  });

// Kota berdasarkan provinsi
export const useCitiesQuery = (provinceId) =>
  useQuery({
    queryKey: ["cities", provinceId],
    queryFn: () => location.getCitiesByProvince(provinceId),
    enabled: !!provinceId, // hanya jalan kalau ada id
    staleTime: 1000 * 60 * 60 * 24,
  });

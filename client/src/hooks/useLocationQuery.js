import { useQuery } from "@tanstack/react-query";
import locationApi from "@/api/location";

// Provinsi
export const useProvincesQuery = () =>
  useQuery({
    queryKey: ["provinces"],
    queryFn: locationApi.getProvinces,
    staleTime: 1000 * 60 * 60 * 24, // cache 1 hari
  });

// Kota berdasarkan provinsi
export const useCitiesQuery = (provinceId) =>
  useQuery({
    queryKey: ["cities", provinceId],
    queryFn: () => locationApi.getCitiesByProvince(provinceId),
    enabled: !!provinceId, // hanya jalan kalau ada id
    staleTime: 1000 * 60 * 60 * 24,
  });

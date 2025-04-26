// hooks/useLocationQuery.js
import loc from "@/services/locations";
import { useQuery } from "@tanstack/react-query";

export const useProvincesQuery = () => {
  return useQuery({
    queryKey: ["provinces"],
    queryFn: () => loc.getProvinces(),
  });
};

export const useSearchProvincesQuery = (query) => {
  return useQuery({
    queryKey: ["searchProvinces", query],
    queryFn: () => loc.searchProvinces(query),
    enabled: !!query, // hanya jalan kalau query ada
  });
};

export const useCitiesByProvinceQuery = (provinceId) => {
  return useQuery({
    queryKey: ["cities", provinceId],
    queryFn: () => loc.getCitiesByProvince(provinceId),
    enabled: !!provinceId, // hanya jalan kalau ada provinceId
  });
};

export const useSearchCitiesQuery = (query) => {
  return useQuery({
    queryKey: ["searchCities", query],
    queryFn: () => loc.searchCities(query),
    enabled: !!query,
  });
};

export const useDistrictsByCityQuery = (cityId) => {
  return useQuery({
    queryKey: ["districts", cityId],
    queryFn: () => loc.getDistrictsByCity(cityId),
    enabled: !!cityId,
  });
};

export const useSubdistrictsByDistrictQuery = (districtId) => {
  return useQuery({
    queryKey: ["subdistricts", districtId],
    queryFn: () => loc.getSubdistrictsByDistrict(districtId),
    enabled: !!districtId,
  });
};

export const usePostalCodesBySubdistrictQuery = (subdistrictId) => {
  return useQuery({
    queryKey: ["postalcodes", subdistrictId],
    queryFn: () => loc.getPostalCodesBySubdistrict(subdistrictId),
    enabled: !!subdistrictId,
  });
};

import { useEffect } from "react";
import { useLocationQuery } from "@/hooks/useLocationQuery";

const SelectLocation = ({ provinceId, cityId, setValue, errors, register }) => {
  const { data: provinces = [], isLoading: loadingProvinces } =
    useLocationQuery("province");
  const { data: cities = [], isLoading: loadingCities } = useLocationQuery(
    "city",
    provinceId
  );

  useEffect(() => {
    const selectedCity = cities.find((c) => c.id.toString() === cityId);
    if (selectedCity) {
      setValue("zipcode", selectedCity.postal_code);
    }
  }, [cityId, cities, setValue]);

  return (
    <div className="flex gap-4">
      {/* Province */}
      <div className="w-1/2">
        <label className="label">Province</label>
        <select
          {...register("province_id")}
          className="input"
          disabled={loadingProvinces}
        >
          <option value="">Select Province</option>
          {provinces.map((p) => (
            <option key={p.id} value={p.id}>
              {p.name}
            </option>
          ))}
        </select>
        {errors.province_id && (
          <p className="text-red-500 text-sm">{errors.province_id.message}</p>
        )}
      </div>

      {/* City */}
      <div className="w-1/2">
        <label className="label">City</label>
        <select
          {...register("city_id")}
          className="input"
          disabled={!provinceId || loadingCities}
        >
          <option value="">Select City</option>
          {cities.map((c) => (
            <option key={c.id} value={c.id}>
              {c.type} {c.name}
            </option>
          ))}
        </select>
        {errors.city_id && (
          <p className="text-red-500 text-sm">{errors.city_id.message}</p>
        )}
      </div>
    </div>
  );
};

export default SelectLocation;

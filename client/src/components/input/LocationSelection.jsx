import {
  useProvincesQuery,
  useCitiesByProvinceQuery,
  useDistrictsByCityQuery,
  useSubdistrictsByDistrictQuery,
  usePostalCodesBySubdistrictQuery,
} from "@/hooks/useLocationQuery";
import { useFormContext, Controller } from "react-hook-form";

const LocationSelection = () => {
  const { control, watch } = useFormContext();

  const selectedCityId = watch("cityId");
  const selectedDistrictId = watch("districtId");
  const selectedProvinceId = watch("provinceId");
  const selectedSubdistrictId = watch("subdistrictId");

  const { data: provinces = [] } = useProvincesQuery();
  const { data: cities = [] } = useCitiesByProvinceQuery(selectedProvinceId);
  const { data: districts = [] } = useDistrictsByCityQuery(selectedCityId);
  const { data: subdistricts = [] } =
    useSubdistrictsByDistrictQuery(selectedDistrictId);
  const { data: postalCodes = [] } = usePostalCodesBySubdistrictQuery(
    selectedSubdistrictId
  );

  return (
    <div className="space-y-4">
      {/* Province */}
      <Controller
        control={control}
        name="provinceId"
        rules={{ required: true }}
        render={({ field }) => (
          <div>
            <label className="block mb-1 font-medium">Province</label>
            <select {...field} className="w-full border p-2 rounded">
              <option value="">Select Province</option>
              {provinces.map((prov) => (
                <option key={prov.id} value={prov.id}>
                  {prov.name}
                </option>
              ))}
            </select>
          </div>
        )}
      />

      {/* City */}
      {selectedProvinceId && (
        <Controller
          control={control}
          name="cityId"
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <label className="block mb-1 font-medium">City</label>
              <select {...field} className="w-full border p-2 rounded">
                <option value="">Select City</option>
                {cities.map((city) => (
                  <option key={city.id} value={city.id}>
                    {city.name}
                  </option>
                ))}
              </select>
            </div>
          )}
        />
      )}

      {/* District */}
      {selectedCityId && (
        <Controller
          control={control}
          name="districtId"
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <label className="block mb-1 font-medium">District</label>
              <select {...field} className="w-full border p-2 rounded">
                <option value="">Select District</option>
                {districts.map((district) => (
                  <option key={district.id} value={district.id}>
                    {district.name}
                  </option>
                ))}
              </select>
            </div>
          )}
        />
      )}

      {/* Subdistrict */}
      {selectedDistrictId && (
        <Controller
          control={control}
          name="subdistrictId"
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <label className="block mb-1 font-medium">Subdistrict</label>
              <select {...field} className="w-full border p-2 rounded">
                <option value="">Select Subdistrict</option>
                {subdistricts.map((subdistrict) => (
                  <option key={subdistrict.id} value={subdistrict.id}>
                    {subdistrict.name}
                  </option>
                ))}
              </select>
            </div>
          )}
        />
      )}

      {/* Postal Code */}
      {selectedSubdistrictId && (
        <Controller
          control={control}
          name="postalCodeId"
          rules={{ required: true }}
          render={({ field }) => (
            <div>
              <label className="block mb-1 font-medium">Postal Code</label>
              <select {...field} className="w-full border p-2 rounded">
                <option value="">Select Postal Code</option>
                {postalCodes.map((postal) => (
                  <option key={postal.id} value={postal.id}>
                    {postal.postalCode}
                  </option>
                ))}
              </select>
            </div>
          )}
        />
      )}
    </div>
  );
};

export default LocationSelection;

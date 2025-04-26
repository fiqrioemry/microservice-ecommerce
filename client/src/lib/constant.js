export const AddressState = {
  name: "",
  address: "",
  province_id: "",
  city_id: z.string().min(2),
  zipcode: z.string().min(5),
  phone: z.string().min(10),
  district: z.string().optional(),
  isMain: z.boolean(),
};

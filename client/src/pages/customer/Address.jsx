import React from "react";

const Address = () => {
  return <div>Address</div>;
};

export default Address;

// import {
//   Dialog,
//   DialogTitle,
//   DialogHeader,
//   DialogContent,
// } from "@/components/ui/dialog";
// import Input from "@/components/ui/input";
// import { useForm } from "react-hook-form";
// import { Button } from "@/components/ui/button";
// import { addressSchema } from "@/lib/schema";
// import { Switch } from "@/components/ui/switch";
// import React, { useState, useEffect } from "react";
// import { Textarea } from "@/components/ui/textarea";
// import { zodResolver } from "@hookform/resolvers/zod";
// import { ScrollArea } from "@/components/ui/scroll-area";
// import { Card, CardContent } from "@/components/ui/card";
// import { useProfileStore } from "@/store/useProfileStore";
// import { useLocationQuery } from "@/hooks/useLocationQuery";
// import { useAddressesQuery } from "@/hooks/useProfileManagement";

// const Address = () => {
//   const { isError, isLoading, data: addresses } = useAddressesQuery();
//   const { addAddress } = useProfileStore();
//   const [dialogOpen, setDialogOpen] = useState(false);
//   const [formTouched, setFormTouched] = useState(false);

//   const {
//     register,
//     handleSubmit,
//     reset,
//     setValue,
//     watch,
//     formState: { errors },
//   } = useForm({
//     resolver: zodResolver(addressSchema),
//     defaultValues: {
//       name: "",
//       address: "",
//       provinceId: "",
//       cityId: "",
//       zipcode: "",
//       phone: "",
//       isMain: false,
//     },
//   });

//   const cityId = watch("city_id");
//   const provinceId = watch("province_id");

//   const { data: provinces = [] } = useLocationQuery("province");
//   const { data: cities = [] } = useLocationQuery("city", provinceId);

//   useEffect(() => {
//     const selectedCity = cities.find((c) => c.id.toString() === cityId);
//     if (selectedCity) {
//       setValue("zipcode", selectedCity.postal_code);
//     }
//   }, [cityId, cities, setValue]);

//   const onSubmit = async (data) => {
//     await addAddress({
//       ...data,
//       province_id: +data.province_id,
//       city_id: +data.city_id,
//     });
//     reset();
//     setFormTouched(false);
//     setDialogOpen(false);
//   };

//   const handleDialogClose = () => {
//     if (formTouched) {
//       const confirmClose = confirm(
//         "Are you sure you want to close? Unsaved data will be lost."
//       );
//       if (!confirmClose) return;
//     }
//     setDialogOpen(false);
//     reset();
//     setFormTouched(false);
//   };

//   if (isLoading) return <div>Loading...</div>;
//   if (isError) return <div>Error. Please try again later.</div>;

//   return (
//     <div className="space-y-4">
//       <div className="flex justify-between items-center">
//         <h2 className="text-lg font-semibold">My Addresses</h2>
//         <Button onClick={() => setDialogOpen(true)}>New Address</Button>
//       </div>

//       <ScrollArea className="h-[500px] space-y-4">
//         {addresses.addresses.map((add) => (
//           <Card key={add.id}>
//             <CardContent className="p-4 space-y-2">
//               <div className="flex items-center justify-between">
//                 <h4 className="text-sm font-medium capitalize">{add.name}</h4>
//                 {add.isMain && (
//                   <span className="text-xs text-white bg-primary rounded px-2">
//                     Main
//                   </span>
//                 )}
//               </div>
//               <p className="text-sm text-muted-foreground">{add.phone}</p>
//               <p className="text-sm">
//                 {add.address}, {add.city}, {add.province}, {add.zipcode}
//               </p>
//               <div className="flex gap-2 text-sm text-primary underline">
//                 <button>Edit</button>
//                 <button>Delete</button>
//               </div>
//             </CardContent>
//           </Card>
//         ))}
//       </ScrollArea>

//       <Dialog open={dialogOpen} onOpenChange={handleDialogClose}>
//         <DialogContent>
//           <DialogHeader>
//             <DialogTitle>Add New Address</DialogTitle>
//           </DialogHeader>
//           <form
//             onSubmit={handleSubmit(onSubmit)}
//             onChange={() => setFormTouched(true)}
//             className="space-y-4"
//           >
//             <Input label="Name" {...register("name")} error={errors.name} />
//             <Textarea
//               label="Address"
//               {...register("address")}
//               error={errors.address}
//             />
//             <div className="flex gap-4">
//               <div className="w-1/2">
//                 <label className="label">Province</label>
//                 <select className="input" {...register("province_id")}>
//                   <option value="">Select Province</option>
//                   {provinces.map((p) => (
//                     <option key={p.id} value={p.id}>
//                       {p.name}
//                     </option>
//                   ))}
//                 </select>
//               </div>
//               <div className="w-1/2">
//                 <label className="label">City</label>
//                 <select className="input" {...register("city_id")}>
//                   <option value="">Select City</option>
//                   {cities.map((c) => (
//                     <option key={c.id} value={c.id}>
//                       {c.type} {c.name}
//                     </option>
//                   ))}
//                 </select>
//               </div>
//             </div>
//             <Input
//               label="Zip Code"
//               {...register("zipcode")}
//               error={errors.zipcode}
//               disabled
//             />
//             <Input
//               label="Phone"
//               type="tel"
//               {...register("phone")}
//               error={errors.phone}
//             />
//             <div className="flex items-center gap-2">
//               <label className="label">Set as Main</label>
//               <Switch {...register("isMain")} />
//             </div>

//             <div className="flex justify-end">
//               <Button type="submit">Save</Button>
//             </div>
//           </form>
//         </DialogContent>
//       </Dialog>
//     </div>
//   );
// };

// export default Address;

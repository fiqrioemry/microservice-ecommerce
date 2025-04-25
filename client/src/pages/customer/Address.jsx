// Updated Address.jsx with new UI/UX and Dialog

import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogContent,
} from "@/components/ui/dialog";
import React, { useState } from "react";
import Input from "@/components/ui/input";
import { useForm } from "react-hook-form";
import { addressSchema } from "@/lib/schema";
import { Switch } from "@/components/ui/switch";
import Button from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Card, CardContent } from "@/components/ui/card";
import { useProfileStore } from "@/store/useProfileStore";

import { useAddressesQuery } from "@/hooks/useProfileManagement";
import { useCitiesQuery, useProvincesQuery } from "@/hooks/useLocationQuery";

const Address = () => {
  const { addAddress } = useProfileStore();
  const [dialogOpen, setDialogOpen] = useState(false);
  const [formTouched, setFormTouched] = useState(false);
  const { isError, isLoading, data: addresses } = useAddressesQuery();

  const {
    register,
    handleSubmit,
    reset,
    watch,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(addressSchema),
    defaultValues: {
      name: "",
      address: "",
      province_id: "",
      city_id: "",
      zipcode: "",
      phone: "",
      isMain: false,
    },
  });

  const provinceId = watch("id");

  const { data: provinces = [] } = useProvincesQuery("province");
  const { data: cities = [] } = useCitiesQuery("city", provinces.id);
  console.log(provinces);
  const onSubmit = async (data) => {
    await addAddress({
      ...data,
      province_id: +data.province_id,
      city_id: +data.city_id,
    });
    reset();
    setFormTouched(false);
    setDialogOpen(false);
  };

  const handleDialogClose = () => {
    if (formTouched) {
      const confirmClose = confirm(
        "Are you sure you want to close? Unsaved data will be lost."
      );
      if (!confirmClose) return;
    }
    setDialogOpen(false);
    reset();
    setFormTouched(false);
  };

  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error. Please try again later.</div>;

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-semibold">My Addresses</h2>
        <Button onClick={() => setDialogOpen(true)}>New Address</Button>
      </div>

      <ScrollArea className="h-[500px] space-y-4">
        {addresses.addresses.map((add) => (
          <Card key={add.id}>
            <CardContent className="p-4 space-y-2">
              <div className="flex items-center justify-between">
                <h4 className="text-sm font-medium capitalize">{add.name}</h4>
                {add.isMain && (
                  <span className="text-xs text-white bg-primary rounded px-2">
                    Main
                  </span>
                )}
              </div>
              <p className="text-sm text-muted-foreground">{add.phone}</p>
              <p className="text-sm">
                {add.address}, {add.city}, {add.province}, {add.zipcode}
              </p>
              <div className="flex gap-2 text-sm text-primary underline">
                <button>Edit</button>
                <button>Delete</button>
              </div>
            </CardContent>
          </Card>
        ))}
      </ScrollArea>

      <Dialog open={dialogOpen} onOpenChange={handleDialogClose}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Add New Address</DialogTitle>
          </DialogHeader>
          <form
            onSubmit={handleSubmit(onSubmit)}
            onChange={() => setFormTouched(true)}
            className="space-y-4"
          >
            <Input label="Name" {...register("name")} error={errors.name} />
            <Textarea
              label="Address"
              {...register("address")}
              error={errors.address}
            />
            <div className="flex gap-4">
              <div className="w-1/2">
                <label className="label">Province</label>
                <select className="input" {...register("province_id")}>
                  {provinces.map((p) => (
                    <option key={p.id} value={p.id}>
                      {p.name}
                    </option>
                  ))}
                </select>
              </div>
              <div className="w-1/2">
                <label className="label">City</label>
                <select className="input" {...register("city_id")}>
                  {cities.map((c) => (
                    <option key={c.id} value={c.id}>
                      {c.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
            <Input
              label="Zip Code"
              {...register("zipcode")}
              error={errors.zipcode}
              disabled
            />
            <Input
              label="Phone"
              type="tel"
              {...register("phone")}
              error={errors.phone}
            />
            <div className="flex items-center gap-2">
              <label className="label">Set as Main</label>
              <Switch {...register("isMain")} />
            </div>
            <div className="flex justify-end">
              <Button type="submit">Save</Button>
            </div>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Address;

---
weight: 10
title: Media Downloads
---

## Media Downloads

* If you want to save files to object storage like S3, you need to perform the corresponding configuration.
* File Download
    * Set Files Requests in Item: In the Item, you need to set Files requests, which include a list of requests for
      downloading files. You can use the `item.SetFilesRequest([]pkg.Request{...})` method to set the list of
      requests.
    * Item.data: Your Item.data field needs to implement a slice of `pkg.File` to store the downloaded file results.
      The name of this field must be "Files," for example: `type DataFile struct { Files []*media.File }`.

      `SetData(&DataFile{})`
    * You can set the fields that are returned. Files []*media.File `json:"files" field:"url,name,ext"`
* Image Download
    * Set Images Requests in Item: In the Item, you need to set Images requests, which include a list of requests
      for downloading images. You can use the `item.SetImagesRequest([]pkg.Request{...})` method to set the list of
      requests.
    * Item.data: Your Item.data field needs to implement a slice of `pkg.Image` to store the downloaded image
      results. The name of this field must be "Images," for
      example: `type DataImage struct { Images []*media.Image }`.

      `SetData(&DataImage{})`
    * You can set the fields that are returned. Images []*media.Image `json:"images" field:"url,name,ext,width,height"`
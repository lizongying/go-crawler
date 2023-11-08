---
weight: 10
title: 媒体下载
---

## 媒体下载

* 如果您希望将文件保存到S3等对象存储中，需要进行相应的配置
* 文件下载
    * 在Item中设置Files请求：在Item中，您需要设置Files请求，即包含要下载的文件的请求列表。
      可以使用`item.SetFilesRequest([]pkg.Request{...})`
      方法设置请求列表。
    * Item.data：您的Item.data字段需要实现pkg.File的切片，用于保存下载文件的结果。
      该字段的名称必须是Files，如`type DataFile struct {Files []*media.File}`。

      `SetData(&DataFile{})`
    * 可以设定返回的字段 Files []*media.File `json:"files" field:"url,name,ext"`
* 图片下载
    * 在Item中设置Images请求：在Item中，您需要设置Images请求，即包含要下载的图片的请求列表。
      可以使用item.SetImagesRequest([]pkg.Request{...})方法设置请求列表。
    * Item.data：您的Item.data字段需要实现pkg.Image的切片，用于保存下载图片的结果。
      该字段的名称必须是Images，如`type DataImage struct {Images []*media.Image}`。

      `SetData(&DataImage{})`
    * 可以设定返回的字段 Images []*media.Image `json:"images" field:"url,name,ext,width,height"`
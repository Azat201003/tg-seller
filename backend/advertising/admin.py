from django.contrib import admin
from advertising.models import Advertise, Filter
from unfold.admin import ModelAdmin

class AdvertiseAdmin(ModelAdmin):
    list_display = ["name", "link"]

admin.site.register(Advertise, AdvertiseAdmin)

class FilterAdmin(ModelAdmin):
    list_display = ["name", "is_active"]
    list_editable = ["is_active"]

admin.site.register(Filter, FilterAdmin)

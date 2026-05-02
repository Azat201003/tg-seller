from advertising.models import Advertise, Filter
from rest_framework import serializers
from advertising import config
import requests
import validators

class AdvertiseSerializer(serializers.ModelSerializer):
    filters = serializers.PrimaryKeyRelatedField(many=True, queryset=Filter.objects.all())

    class Meta:
        model = Advertise
        fields = ["advertise_id", "link", "name", "filters"]
    
    def validate_link(self, value):
        if config.VALIDATE_LINK_AVAILABILITY:
            if requests.get(value).status_code != 200:
                raise serializers.ValidationError("Url should be available (200 status code)")
            return value
        validation = validators.url(value)
        if validation != True:
            raise serializers.ValidationError("Url isn't valid")
        return value

class FilterSerializer(serializers.ModelSerializer):
    class Meta:
        model = Filter
        fields = ["filter_id", "is_active", "name"]

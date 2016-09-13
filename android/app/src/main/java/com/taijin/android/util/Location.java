package com.taijin.android.util;

/**
 * Created by taijin on 9/3/16.
 */
public class Location {
    private double lat;
    private double lng;

    public double getLat() {
        return lat;
    }

    public void setLat(double lat) {
        this.lat = lat;
    }

    public double getLng() {
        return lng;
    }

    public void setLng(double lng) {
        this.lng = lng;
    }

    public Location(double lat, double lng) {

        this.lat = lat;
        this.lng = lng;
    }
}

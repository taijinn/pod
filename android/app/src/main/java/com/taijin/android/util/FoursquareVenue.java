package com.taijin.android.util;

import com.google.gson.annotations.SerializedName;

/**
 * Created by taijin on 9/4/16.
 */
public class FoursquareVenue {
    @SerializedName("name")
    public String name;
    @SerializedName("city")
    public String city;
    //@SerializedName("id")
    //private String category;

    public FoursquareVenue(String name, String city) {

        this.name = name;
        this.city = city;
        //this.category = category;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }
}

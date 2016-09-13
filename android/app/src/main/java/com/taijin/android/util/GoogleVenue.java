package com.taijin.android.util;

import com.google.gson.annotations.SerializedName;

/**
 * Created by taijin on 9/6/16.
 */
public class GoogleVenue {
    @SerializedName("name")
    private String name;
    @SerializedName("rating")
    private double rating;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public double getRating() {
        return rating;
    }

    public void setRating(double rating) {
        this.rating = rating;
    }
}

package com.taijin.android.util;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by taijin on 9/6/16.
 */
public class GoogleVenueResult {
    @SerializedName("results")
    private List<GoogleVenue> venues;

    public List<GoogleVenue> getVenues() {
        return venues;
    }

    public void setVenues(List<GoogleVenue> venues) {
        this.venues = venues;
    }
}

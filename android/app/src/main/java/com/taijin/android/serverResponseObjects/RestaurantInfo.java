package com.taijin.android.serverResponseObjects;

/**
 * Created by taijin on 8/17/16.
 */
public class RestaurantInfo {

    String ID;
    String Name;
    String Phone;

    @Override
    public String toString() {
        return Name + " (" + Phone + ")";
    }
}

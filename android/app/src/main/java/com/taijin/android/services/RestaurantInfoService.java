package com.taijin.android.services;

import com.taijin.android.dataRetrievingObjects.DataGettingResInfo;
import com.taijin.android.serverResponseObjects.RestaurantInfo;

import java.util.List;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;

/**
 * Created by taijin on 8/17/16.
 */
public interface RestaurantInfoService {
    //@GET("https://maps.googleapis.com/maps/api/place/nearbysearch/json?key={key}&location={longtitude},{latitude}&radius={radius}&type={type}&opennow={opennow}")
    @GET("")
    Call<List<RestaurantInfo>> getRestaurantInfo(@Path("key") String key,
                                                @Path("longtitude") double longtitude,
                                                @Path("latitude") double latitude,
                                                @Path("radius") double radius,
                                                @Path("type") String type,
                                                @Path("opennow") boolean opennow);

    @POST("{owner}")
        //Call<List<RestaurantInfo>> repoContributors(@Body DataGettingResInfo data,
        //    @Path("owner") String owner);
    Call<RestaurantInfo> getRestaurantInfo(@Body DataGettingResInfo data,
                                           @Path("owner") String owner);
    //     @Path("repo") String repo);


    public static final Retrofit retrofit = new Retrofit.Builder()
            .baseUrl("http://10.0.2.2:8080/")
            .addConverterFactory(GsonConverterFactory.create())
            .build();
}



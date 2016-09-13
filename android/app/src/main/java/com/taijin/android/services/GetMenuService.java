package com.taijin.android.services;

import com.taijin.android.dataRetrievingObjects.DataGettingResInfo;
import com.taijin.android.serverResponseObjects.RestaurantInfo;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.Body;
import retrofit2.http.POST;
import retrofit2.http.Path;

/**
 * Created by taijin on 8/29/16.
 */
public interface GetMenuService {
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

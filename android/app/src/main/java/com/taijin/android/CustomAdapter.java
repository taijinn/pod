package com.taijin.android;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import com.taijin.android.util.GoogleVenue;

import java.util.List;

/**
 * Created by taijin on 9/2/16.
 */
public class CustomAdapter extends ArrayAdapter<GoogleVenue> {

    private Context context;
    private List<GoogleVenue> venueList;

    public CustomAdapter(Context context, List<GoogleVenue> resource) {
        super(context, R.layout.custom_row,resource);
        this.context = context;
        this.venueList = resource;

    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        LayoutInflater dashenInflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View view = dashenInflater.inflate(R.layout.custom_row, parent, false);
        GoogleVenue singleFoodItem = getItem(position);
        TextView text = (TextView) view.findViewById(R.id.listText);
        text.setText(singleFoodItem.getName());
        //ImageView image = (ImageView) view.findViewById(R.id.listImage);

        //text.setText(singleFoodItem.getName()+", "+singleFoodItem.getCity()+", "+singleFoodItem.getCategory());
        //image.setImageResource(R.drawable.test);
        return view;
    }
}

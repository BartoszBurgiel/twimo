// Import core functionality dependencies
import React from "react";
import { View, Text, StyleSheet, Image, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import the color utility
import color from "../../utils/colorPallet";

// Define the card component
const Card = ({ data }) => {
  // Function to render star icons based on rating value
  const ratingCounter = () => {
    // Set up targetArray
    let arrayOfIcons = [];
    // Loop ratingValue times
    for (let i = 0; i < 5; i++) {
      if (i < data.rating) {
        arrayOfIcons.push(
          <Ionicons name="md-star" key={i} size={25} color={color.brand.dark} />
        );
      } else {
        arrayOfIcons.push(
          <Ionicons
            name="md-star-outline"
            key={i}
            size={25}
            color={color.brand.dark}
          />
        );
      }
    }
    //Return text element containing the array
    return <Text style={styles.cardMetrics}>{arrayOfIcons}</Text>;
  };

  return (
    <View style={styles.card}>
      <View style={styles.pictureWrapper}>
        <Image
          style={styles.image}
          source={{ uri: "https://placekitten.com/180/180" }}
        ></Image>
      </View>
      <View style={styles.textWrapper}>
        <Text style={styles.cardHeader}>{data.name}</Text>
        <Text style={styles.cardSubHeader}>{data.location}</Text>
        {ratingCounter()}
        <Text style={styles.cardMetrics}>
          <Ionicons name="logo-euro" size={20} color={color.brand.dark} />{" "}
          {data.pricing} / 5
        </Text>
        <TouchableOpacity style={styles.button}>
          <Text style={styles.buttonText}>Mehr Infos</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
};

// Define the stylesheet
const styles = StyleSheet.create({
  card: {
    backgroundColor: color.white,
    width: 350,
    height: 200,
    padding: 10,
    marginTop: 20,
    marginBottom: 20,
    shadowColor: color.gray.darker,
    shadowOpacity: 0.25,
    borderRadius: 12,
    shadowOffset: {
      width: 0,
      height: 5
    },
    shadowRadius: 13,
    flexDirection: "row"
  },
  image: {
    height: 180,
    width: 180,
    borderRadius: 12
  },
  textWrapper: {
    justifyContent: "space-around",
    paddingLeft: 10,
    width: 150
  },
  cardHeader: {
    fontWeight: "900",
    fontSize: 25,
    color: color.brand.normal
  },
  cardSubHeader: {
    fontWeight: "300",
    textTransform: "uppercase",
    fontSize: 12
  },
  cardMetrics: {
    fontWeight: "600"
  },
  button: {
    backgroundColor: color.brand.light,
    height: 30,
    width: 120,
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 20
  },
  buttonText: {
    textTransform: "uppercase",
    color: color.brand.dark,
    fontWeight: "800"
  }
});

// Export the Card component to be used inside the App
export default Card;

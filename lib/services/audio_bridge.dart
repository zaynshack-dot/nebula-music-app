import 'dart:convert';
import 'package:crypto/crypto.dart';
import 'package:just_audio/just_audio.dart';

class OmniStreamAudioBridge {
  static const String _proxyUrl = 'http://127.0.0.1:8585/stream';
  static const String _secretKey = 'OmniStream_HMAC_Super_Secret_Key_2026_!#';
  final AudioPlayer player = AudioPlayer();

  String _sign(String trackId, String ts) {
    final hmacSha256 = Hmac(sha256, utf8.encode(_secretKey));
    return hmacSha256.convert(utf8.encode('$trackId:$ts')).toString();
  }

  Future<void> streamTrack(String id, String srcUrl, String tier) async {
    final ts = DateTime.now().millisecondsSinceEpoch.toString();
    final proxyUri = Uri.parse(_proxyUrl).replace(queryParameters: {
      'id': id, 'src': srcUrl, 'ts': ts, 'sig': _sign(id, ts), 'q': tier,
    });
    await player.setUrl(proxyUri.toString());
    player.play();
  }
}